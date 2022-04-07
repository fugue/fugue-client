package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/textproto"
	"path/filepath"
	"strings"

	"github.com/fugue/fugue-client/client/custom_rules"
	"github.com/fugue/fugue-client/client/families"
	"github.com/fugue/fugue-client/models"
	"github.com/fugue/regula/v2/pkg/regotools/metadoc"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	targetRuleFamilies string
)

type regoFile struct {
	Name         string
	Providers    []string
	ResourceType string
	Description  string
	Severity     string
	Text         string
}

func (rego *regoFile) ParseText() error {
	md, err := metadoc.RegoMetaFromString(rego.Text)
	if err != nil {
		return err
	}

	if md.HasMetadoc() {
		rego.Name = md.Title
		rego.Providers = md.Providers
		rego.ResourceType = md.ResourceType
		rego.Description = md.Description
		rego.Severity = md.Severity
		if rego.Name == "" {
			return errors.New("expected a 'title' entry in the __rego__metadoc__")
		}
		if rego.ResourceType == "" {
			return errors.New("expected a resource_type in the rule")
		}
		if rego.Description == "" {
			return errors.New("expected a 'description' entry in the __rego__metadoc__")
		}
		if len(rego.Providers) == 0 {
			return errors.New("expected a 'custom.providers' entry in the __rego__metadoc__")
		}
		if rego.Severity == "" {
			rego.Severity = "High"
		}
	} else {
		// Extract the header block from the first comment block.
		headerText := ""
		inFirstCommentBlock := false
		for _, line := range strings.Split(rego.Text, "\n") {
			if strings.HasPrefix(line, "#") {
				if !inFirstCommentBlock && headerText == "" {
					inFirstCommentBlock = true
				}
				if inFirstCommentBlock {
					headerText += strings.TrimSpace(strings.TrimPrefix(line, "#"))
					headerText += "\r\n"
				}
			} else {
				inFirstCommentBlock = false
			}
		}

		// Parse the HTTP headers in `headerText`.
		reader := bufio.NewReader(strings.NewReader(headerText + "\r\n"))
		tp := textproto.NewReader(reader)
		headers := make(map[string][]string)
		mimeHeader, err := tp.ReadMIMEHeader()
		if err != nil {
			log.Fatalf("Could not find a metadoc entry or HTTP style headers: %s", err)
		} else {
			headers = http.Header(mimeHeader)
		}

		// Helper to obtain a specific header.
		getHeader := func(name string) string {
			if arr, ok := headers[name]; ok {
				if len(arr) > 0 {
					return arr[0]
				}
			}
			return ""
		}

		if provider := getHeader("Provider"); provider != "" {
			rego.Providers = []string{provider}
		}
		rego.ResourceType = getHeader("Resource-Type")
		rego.Description = getHeader("Description")
		rego.Severity = getHeader("Severity")

		if rego.ResourceType == "" {
			return errors.New("expected a resource type by the header \"Resource-Type\"")
		}
		if rego.Description == "" {
			return errors.New("expected a description by the header \"Description\"")
		}
		if len(rego.Providers) == 0 {
			return errors.New("expected a provider by the header \"Provider\"")
		}
		if rego.Severity == "" {
			rego.Severity = "High"
		}
	}

	return nil
}

func loadRego(path string) (*regoFile, error) {

	baseName := filepath.Base(path)
	extension := strings.ToLower(filepath.Ext(path))
	if extension != ".rego" {
		return nil, nil
	}

	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if len(contents) == 0 {
		return nil, nil
	}

	name := baseName[:len(baseName)-len(extension)]
	name = strings.ReplaceAll(name, "_", "-")

	rego := regoFile{
		Text: string(contents),
		Name: name,
	}

	err = rego.ParseText()

	return &rego, err
}

var familyToUUIDCache map[string]string = nil

func familyToUUID(family string) (string, bool) {
	if familyToUUIDCache == nil {
		client, auth := getClient()
		familyToUUIDCache = make(map[string]string)

		isTruncated := true
		currentOffset := int64(0)
		maxItems := int64(10)
		for isTruncated {
			params := families.NewListFamiliesParams()
			params.Offset = &currentOffset
			params.MaxItems = &maxItems

			if famList, err := client.Families.ListFamilies(params, auth); err == nil {
				for _, famInfo := range famList.Payload.Items {
					if !(famInfo.Source == "CUSTOM" || famInfo.ID == "Custom") {
						continue
					}
					familyToUUIDCache[famInfo.Name] = famInfo.ID
				}

				isTruncated = famList.Payload.IsTruncated
				currentOffset = famList.Payload.NextOffset
			} else {
				logrus.Fatalf("Error fetching families: %s\n", err)
			}
		}
	}

	if uuid, ok := familyToUUIDCache[family]; ok {
		return uuid, true
	} else {
		return "", false
	}
}

// NewSyncRulesCommand returns a command that watches a directory for changes
// to rego files
func NewSyncRulesCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "rules [directory]",
		Short: "Sync rules to the organization",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			client, auth := getClient()

			getRuleByName := func(name string) *models.CustomRule {
				params := custom_rules.NewListCustomRulesParams()
				resp, err := client.CustomRules.ListCustomRules(params, auth)
				CheckErr(err)
				for _, rule := range resp.Payload.Items {
					if rule.Name == name {
						return rule
					}
				}
				return nil
			}

			var targetFamilies []string = nil
			if targetRuleFamilies != "" {
				targetFamilies = strings.Split(targetRuleFamilies, ",")
			}

			updateRule := func(path string) error {

				rego, err := loadRego(path)
				if err != nil {
					CheckErr(err)
					return nil
				}
				if rego == nil {
					return nil
				}

				var ruleFamilies []string = nil

				// We want to allow people to override the families specified in the
				// __rego__metadoc__ block using the command line.
				if len(targetFamilies) == 0 {
					md, err := metadoc.RegoMetaFromPath(path)
					if err != nil {
						log.Fatal(err)
					}
					if len(md.Families) != 0 {
						for _, family := range md.Families {
							if _, err := uuid.Parse(family); err == nil {
								ruleFamilies = append(ruleFamilies, family)
							} else {
								if familyUUID, ok := familyToUUID(family); !ok {
									log.Fatalf("Unable to find family '%s' (referenced in '%s').", family, path)
								} else {
									ruleFamilies = append(ruleFamilies, familyUUID)
								}
							}
						}
					} else {
						ruleFamilies = make([]string, 0)
					}
				} else {
					ruleFamilies = targetFamilies
				}

				existingRule := getRuleByName(rego.Name)
				if existingRule == nil {
					fmt.Println("Creating rule", rego.Name)
					params := custom_rules.NewCreateCustomRuleParams()
					params.Rule = &models.CreateCustomRuleInput{
						Name:         rego.Name,
						Description:  rego.Description,
						ResourceType: rego.ResourceType,
						Providers:    rego.Providers,
						Severity:     rego.Severity,
						RuleText:     rego.Text,
					}

					created, err := client.CustomRules.CreateCustomRule(params, auth)
					if err != nil {
						log.Fatal(err)
					}

					if ruleFamilies != nil {
						famParams := custom_rules.NewUpdateCustomRuleParams()
						famParams.RuleID = created.Payload.ID
						famParams.Rule = &models.UpdateCustomRuleInput{
							Families: ruleFamilies,
						}
						if _, err := client.CustomRules.UpdateCustomRule(famParams, auth); err != nil {
							log.Fatal(err)
						}
					}
				} else {
					fmt.Println("Updating rule", rego.Name)
					params := custom_rules.NewUpdateCustomRuleParams()
					params.RuleID = existingRule.ID
					params.Rule = &models.UpdateCustomRuleInput{
						Description:  rego.Description,
						ResourceType: rego.ResourceType,
						Providers:    rego.Providers,
						Severity:     rego.Severity,
						RuleText:     rego.Text,
						Families:     ruleFamilies,
					}
					if _, err := client.CustomRules.UpdateCustomRule(params, auth); err != nil {
						log.Fatal(err)
					}
				}
				return nil
			}

			files, err := ioutil.ReadDir(args[0])
			CheckErr(err)

			for _, file := range files {
				if filepath.Ext(file.Name()) == ".rego" {
					updateRule(filepath.Join(args[0], file.Name()))
				}
			}
		},
	}
	cmd.Flags().StringVarP(&targetRuleFamilies, "target-rule-family", "", "", "Comma separated list of UUID of families")
	return cmd
}

func init() {
	syncCmd.AddCommand(NewSyncRulesCommand())
}
