package main

import (
    "./github"
    "flag"
    "log"
    "os"
)

var verbose = flag.Bool("v", false, "report issues in verbose mode")

func init() {
    flag.Parse()
}

func main() {
    i := 0
    args := make([]string, len(os.Args[1:]))
    for _, arg := range os.Args[1:] {
        if arg == "-v" { continue }
        args[i] = arg
        i++
    }
    args = args[:i]
    result, err := github.SearchIssues(args)
    if err != nil {
       log.Fatal(err)
    }
    if *verbose {
        result.FormattedReport()
    } else {
        result.SimpleReport()
    }
}