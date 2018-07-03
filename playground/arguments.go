package main

import (
	"flag"
)

func main() {
	numPtr := flag.Int("n",1,"number of greetings")
	strPtr := flag.String("name","username","name of user")
	titlePtr := flag.Bool("title", false,"print title before greeting")
	flag.Parse()

	if numPtr != nil {
		println(*numPtr)
	}

	if strPtr != nil {
		println(*strPtr)
	}

	if titlePtr != nil {
		println(*titlePtr)
	}
}