package main

import (
	"fmt"
	"os"
	"strings"
)

const usage = `
date-delta <from-time> [- <to-time>] [in <unit>]

Prints the delta between two dates. If <to-time> is not
given, "now" is assumed.`

func main() {

	args := os.Args

	if len(args) <= 1 || (len(args) == 2 && args[1] == "-h") {
		fmt.Println(usage)
		return
	}

	args = args[1:]
	tA, tB, unit, err := parseCmdLine(strings.Join(args, " "))

	if err != nil {

		fmt.Println("error: ", err)
		os.Exit(1)
	}

	duration := tA.Sub(tB).Abs()

	fmt.Println(unit.durationAsString(duration))
}

