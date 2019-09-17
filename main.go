package main

import "github.com/fugue/fugue-client/cmd"

var (
	version = "localVersion"
	commit  = "localCommit"
)

func main() {
	cmd.Execute(version, commit)
}
