package main

import (
	"os"
	"net/http"
	"fmt"
	"io/ioutil"
	"flag"
	"strings"
)

var AppName string

func main() {
	AppName = os.Args[0]

	fileNamePtr := flag.String("hosts", "hosts.txt", "List of URLs")
	flag.Parse()

	data, err := ioutil.ReadFile(*fileNamePtr)
	exitIfError(err)

	hosts := strings.TrimRight(string(data), "\n")

	for _, host := range strings.Split(hosts, "\n") {
		fmt.Printf("Requesting host: %s\n", host)
		resp, err := http.Get(host)
		exitIfError(err)
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		exitIfError(err)
		fmt.Printf("%s", b)
		fmt.Println()
	}
}

func exitIfError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", AppName, err)
		os.Exit(1)
	}
}