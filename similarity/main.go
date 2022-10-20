package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Run with -help for help information")
	}

	var in string
	var help bool
	flag.StringVar(&in, "in", "sentence_files/tiny.txt", "the input file path")
	flag.BoolVar(&help, "help", false, "use to display help text")
	flag.Parse()

	if help {
		pad := func() {
			fmt.Printf("\n\n")
		}
		pad()
		fmt.Println(" Welcome to the sentence similarity tool!")
		fmt.Print(" This code is licensed under GPLv3")
		pad()
		fmt.Println(" Args:")
		fmt.Println("  -in       The file name for password input. Defaults to pass/common.txt")
		fmt.Println("  -help     Display this help text :)")
		pad()
		return
	}

	driver(in)
}

func driver(inFile string) {
	fmt.Println(inFile)
}
