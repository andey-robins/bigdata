package main

import (
	"flag"
	"fmt"

	"github.com/andey-robins/bigdata/similarity/sentences"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Run with -help for help information")
	}

	var in string
	var help bool
	var k int
	flag.StringVar(&in, "in", "sentence_files/tiny.txt", "the input file path")
	flag.IntVar(&k, "k", 0, "the distance count")
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
		fmt.Println("  -in       The file name for password input. Defaults to sentence_files/tiny.txt")
		fmt.Println("  -k        The distance measure. Defaults to 0")
		fmt.Println("  -help     Display this help text :)")
		pad()
		return
	}

	if k == 0 {
		driver_0(in)
	} else {
		driver(in)
	}
}

func driver(inFile string) {
	fmt.Println(inFile)
}

func driver_0(inFile string) {
	ss := sentences.New()
	ss.LoadFile(inFile)
	count := ss.CountDupes()
	fmt.Printf("File '%s' has %v duplicate lines.\n", inFile, count)
}
