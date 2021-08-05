//lint:file-ignore U1000 Ignore all unused code
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/fugue/fugue-client/client"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/spf13/cobra"
)

const (
	// DefaultHost is the default hostname of the Fugue API
	DefaultHost = "api.riskmanager.fugue.co"

	// DefaultBase is the base path of the Fugue API
	DefaultBase = "v0"
)

func mustGetEnv(name string) string {
	value := os.Getenv(name)
	if value == "" {
		fmt.Fprintf(os.Stderr, "Missing environment variable: %s\n", name)
		os.Exit(1)
	}
	return value
}

func getEnvWithDefault(name, defaultValue string) string {
	value := os.Getenv(name)
	if value == "" {
		return defaultValue
	}
	return value
}

func getClient() (*client.Fugue, runtime.ClientAuthInfoWriter) {

	clientID := mustGetEnv("FUGUE_API_ID")
	clientSecret := mustGetEnv("FUGUE_API_SECRET")

	host := getEnvWithDefault("FUGUE_API_HOST", DefaultHost)
	base := getEnvWithDefault("FUGUE_API_BASE", DefaultBase)

	transport := httptransport.New(host, base, []string{"https"})
	apiclient := client.New(transport, strfmt.Default)

	auth := httptransport.BasicAuth(clientID, clientSecret)

	return apiclient, auth
}

func showResponse(obj interface{}) {
	js, err := json.Marshal(obj)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", string(js))
}

func flagStringValue(cmd *cobra.Command, name string) string {
	value, err := cmd.PersistentFlags().GetString(name)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

func flagBoolValue(cmd *cobra.Command, name string) bool {
	value, err := cmd.PersistentFlags().GetBool(name)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

func flagInt64Value(cmd *cobra.Command, name string) int64 {
	value, err := cmd.PersistentFlags().GetInt64(name)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

func flagStringSliceValue(cmd *cobra.Command, name string) []string {
	if cmd.PersistentFlags().Lookup(name).Changed {
		value, err := cmd.PersistentFlags().GetStringSlice(name)
		if err != nil {
			log.Fatal(err)
		}
		return value
	}
	return nil
}
