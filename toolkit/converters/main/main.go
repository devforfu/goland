package main

import (
	"flag"
	"io/ioutil"
	"os"
	"fmt"
	"strings"
	"strconv"
	"../temperature"
)

var path = flag.String("path", "", "path to file with measurements")

type Temperature struct {
	amount float64
	system string
}

func main() {
	flag.Parse()
	temperatures := parseTemperatures(*path)

	var celsius []temperature.Celsius
	var fahrenheit []temperature.Fahrenheit

	for _, temp := range temperatures {
		switch temp.system {
		case "C": celsius = append(celsius, temperature.Celsius(temp.amount))
		case "F": fahrenheit = append(fahrenheit, temperature.Fahrenheit(temp.amount))
		}
	}

	println("Celsius measurements:")
	for _, c := range celsius {
		println(c.String())
	}

	println("Fahrenheit measurements:")
	for _, f := range fahrenheit {
		println(f.String())
	}
}

func parseTemperatures(path string) []Temperature {
	data, err := ioutil.ReadFile(path)
	exitOnError(err)
	var pairs []Temperature
	for index, string := range strings.Split(string(data), "\n") {
		if string == "" {
			fmt.Printf("Empty string at line %d", index)
			continue
		}
		strings := strings.Split(string, ",")
		amountString, system := strings[0], strings[1]
		value, err := strconv.ParseFloat(amountString, 64)
		if err != nil {
			continue
		}
		pairs = append(pairs, Temperature{value, system})
	}
	return pairs
}

func exitOnError(err error) {
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}

