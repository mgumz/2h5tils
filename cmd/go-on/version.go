package main

import "fmt"

var (
	version   = "0.9.0"
	gitHash   = ""
	buildDate = ""
)

func printVersion() {

	fmt.Println("goon:\t" + version)
	if gitHash != "" {
		fmt.Println("git:\t" + gitHash)
	}
	if buildDate != "" {
		fmt.Println("build:\t" + buildDate)
	}
}
