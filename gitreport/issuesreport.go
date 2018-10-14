package main

import (
    "./github"
    "log"
    "os"
)

func main() {
    result, err := github.SearchIssues(os.Args[1:])
    if err != nil {
        log.Fatal(err)
    }
    result.FormattedReport()
}