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
var captureJSON bool

func isOutputJSON() bool {
	// We need to to os Args because rootCmd.Execute() has not run yet
	argsWithoutProg := os.Args[1:]
	for _, elem := range argsWithoutProg {
		if elem == "--json" {
			return true
		}
	}
	return false
}

func printLastJSONFromTheOutput(out []byte) {
	outAsString := string(out)
	doubleEnterRegex := regexp.MustCompile(`\n\s\n`)
	outAsArray := doubleEnterRegex.Split(outAsString, -1)
	src := []byte(outAsArray[len(outAsArray)-1])
	dst := &bytes.Buffer{}
	if err := json.Indent(dst, src, "", "  "); err != nil {
		panic(err)
	}
	fmt.Println(dst.String())
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

	if captureJSON {
		wStdout.Close()
		wStderr.Close()
		err, _ := ioutil.ReadAll(rStderr)
		os.Stdout = rescueStdout
		os.Stderr = rescueStderr
		printLastJSONFromTheOutput(err)
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
