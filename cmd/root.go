package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
var captureJSON bool

//jsonPositionToShow When assigned in the children commands gets the nth Json in the output
var jsonPositionToShow int = -1

func isOutputJSON() bool {
	// We need to to os Args because rootCmd.Execute() has not run yet
	argsWithoutProg := os.Args[1:]
	for _, elem := range argsWithoutProg {
		if strings.ToLower(elem) == "--json" {
			return true
		}
	}
	return false
}

func printJSONOutput(out []byte) {

	outArray := bytes.Split(out, []byte("\n"))

	var jsonArray []string
	for _, v := range outArray {
		var js map[string]interface{}
		if json.Unmarshal(v, &js) == nil {
			jsonArray = append(jsonArray, string(v))
		}
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
	fmt.Println(elemToPrintIndented.String())
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fugue",
	Short: "Fugue API Client",
	Long:  ``,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version, commit string) {
	rootCmd.Version = fmt.Sprintf("%s-%s", version, commit)

	var rescueStdout, wStdout *os.File
	var rescueStderr, rStderr, wStderr *os.File
	if isOutputJSON() {
		os.Setenv("DEBUG", "1")
		rescueStdout, rescueStderr = os.Stdout, os.Stderr
		_, wStdout, _ = os.Pipe()
		rStderr, wStderr, _ = os.Pipe()
		os.Stdout, os.Stderr = wStdout, wStderr
	}

	CheckErr(rootCmd.Execute())

	if isOutputJSON() && captureJSON {
		wStdout.Close()
		wStderr.Close()
		err, _ := ioutil.ReadAll(rStderr)
		os.Stdout = rescueStdout
		os.Stderr = rescueStderr
		printJSONOutput(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Application global flags
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.fugue.yaml)")
	rootCmd.PersistentFlags().BoolVar(&captureJSON, "json", false, "outputs the Fugue API JSON response")
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
		// add newline if needed
		if !strings.HasSuffix(msg, "\n") {
			msg += "\n"
		}
		fmt.Fprint(os.Stderr, msg)
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
