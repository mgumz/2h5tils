package main

import "fmt"

const (

	usage string = `usage: goon <SPEC>

SPEC:

@hour    - run at next full hour
@minute  - run at next full minute
@tens    - run at the next full 10 minutes
@quarter - run at the next quarter

:15      - continue at <current_hour>:15
20:15    - go on at 20:15
11:15PM  - go on at 11:15PM

30s      - go on in 30s (aka "sleep 30")
2m       - go on in 2 minutes aka 120s (aka "sleep 120")`
)

func printUsage() {
	fmt.Println(usage)
}
