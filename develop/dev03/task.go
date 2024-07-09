package main

import "flag"

func main() {
	inputFile := flag.String("input", "", "Input file")
	outputFile := flag.String("output", "", "Output file")

	keys := make(map[string]string)
	flag.StringVar(&keys["-k"], "k", "", "Column for sorting")
	flag.StringVar(&keys["-n"], "n", "", "Numeric sort")
	flag.StringVar(&keys["-r"], "r", "", "Reverse order")
	flag.StringVar(&keys["-u"], "u", "", "Unique lines")
}
