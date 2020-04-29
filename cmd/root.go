package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	// DefaultErrorExitCode is the exit code value when an error occurs
	DefaultErrorExitCode = 1
)

var cfgFile string
var outputFormat string

//jsonPositionToShow When assigned in the children commands gets the nth Json in the output
var jsonPositionToShow int = -1

func isOutputJSON() bool {
	// We need to check the command line Args because rootCmd.Execute() has not run yet
	argsWithoutProg := os.Args[1:]
	matched, _ := regexp.MatchString(`(?i)--output json`, strings.Join(argsWithoutProg, " "))
	return matched
}

func jsonOutput(out []byte) string {
	outArray := bytes.Split(out, []byte("\n"))
	var jsonArray []string
	for _, v := range outArray {
		var js map[string]interface{}
		if json.Unmarshal(v, &js) == nil {
			jsonArray = append(jsonArray, string(v))
		}
	}

	if len(jsonArray) == 0 {
		return ""
	}

	var elemToPrint string
	if jsonPositionToShow == -1 {
		elemToPrint = jsonArray[len(jsonArray)-1]
	} else {
		elemToPrint = jsonArray[jsonPositionToShow]
	}

	elemToPrintIndented := &bytes.Buffer{}
	if err := json.Indent(elemToPrintIndented, []byte(elemToPrint), "", "  "); err != nil {
		panic(err)
	}
	return elemToPrintIndented.String()
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fugue",
	Short: "Fugue API Client",
	Long:  ``,
	// This is a hack to check the flag output is valid.
	// wait for this to be merged: https://github.com/spf13/pflag/issues/236
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		validOutputs := []string{"table", "json"}
		valid := func() bool {
			for _, validOutput := range validOutputs {
				if strings.ToLower(outputFormat) == validOutput {
					return true
				}
			}
			return false
		}
		if !valid() {
			Fatal(fmt.Sprintf("Value '%s' is invalid for flag 'output'. Valid values are: %v",
				outputFormat, validOutputs), DefaultErrorExitCode)
		}
	},

	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

var rescueStdout, rStdout, wStdout *os.File
var rescueStderr, rStderr, wStderr *os.File

func getRedirectedOutput() ([]byte, []byte) {
	wStdout.Close()
	wStderr.Close()
	out, _ := ioutil.ReadAll(rStdout)
	err, _ := ioutil.ReadAll(rStderr)
	os.Stdout = rescueStdout
	os.Stderr = rescueStderr
	return out, err
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version, commit string) {
	rootCmd.Version = fmt.Sprintf("%s-%s", version, commit)

	rootCmd.SetOutput(os.Stderr)

	if isOutputJSON() {
		os.Setenv("DEBUG", "1")
		rescueStdout, rescueStderr = os.Stdout, os.Stderr
		rStdout, wStdout, _ = os.Pipe()
		rStderr, wStderr, _ = os.Pipe()
		os.Stdout, os.Stderr = wStdout, wStderr
	}

	CheckErr(rootCmd.Execute())

	if isOutputJSON() {
		_, err := getRedirectedOutput()
		outStr := jsonOutput(err)
		fmt.Println(outStr)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Application global flags
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.fugue.yaml)")
	rootCmd.PersistentFlags().StringVar(&outputFormat, "output", "table", "Use 'json' to access the Fugue API JSON response")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		CheckErr(err)

		// Search config in home directory with name ".fugue" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".fugue")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// Fatal prints the message (if provided) and then exits.
func Fatal(msg string, code int) {
	if len(msg) > 0 {

		if isOutputJSON() {
			_, err := getRedirectedOutput()
			outStr := jsonOutput(err)

			if outStr != "" {
				fmt.Fprintln(os.Stderr, outStr)
			} else {
				fmt.Fprintf(os.Stderr, `{"error": "%s"}\\n`, msg)
			}
		} else {
			// add newline if needed
			if !strings.HasSuffix(msg, "\n") {
				msg += "\n"
			}
			fmt.Fprint(os.Stderr, msg)
		}

	}
	os.Exit(code)
}

// CheckErr prints a user friendly error to STDERR and exits with a non-zero
// exit code. Unrecognized errors will be printed with an "error: " prefix.
func CheckErr(err error) {
	if err == nil {
		return
	}
	msg := err.Error()
	if !strings.HasPrefix(msg, "error: ") {
		msg = fmt.Sprintf("error: %s", msg)
	}
	Fatal(msg, DefaultErrorExitCode)
}
