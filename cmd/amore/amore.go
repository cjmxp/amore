// The amore command for bundling assets and helpers for quick project creation.
package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	amoreVersion = "0.0.5" // command line version
)

var (
	namePackage    = flag.String("pkg", "main", "name of the go package for the source to be generated in")
	nameSourceFile = flag.String("out", "asset_bundle.go", "name of the outputted file for bundling")
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) > 0 {
		if args[0] == "bundle" {
			bundle(args[1:]...)
		} else if args[0] == "init" || args[0] == "new" {
			projectName := ""
			if len(args) > 1 {
				projectName = args[1]
			}
			newProject(projectName)
		} else if args[0] == "version" {
			printVersion()
		} else {
			printHelp()
		}
	} else {
		printHelp()
	}
}

// printVersion outputs the command line version
func printVersion() {
	fmt.Printf("Amore version: %v", amoreVersion)
}

func printHelp() {
	fmt.Printf(`
Amore - %v

usage: amore <command> [args ..]

Commands:
  new <project_name>
    Generates all basic files and folders to get started in the current directory

  bundle [paths...]
    Bundles inputed file paths.
    If no paths were given then it will bundle the default assets folder and the conf.toml file.

  version
    Will print the version of this cli.

  help
    Will print this message.

`, amoreVersion)
}

// Prints out the error message and exists with a non-success signal.
func exitWithError(err error) {
	fmt.Println(err)
	os.Exit(1)
}
