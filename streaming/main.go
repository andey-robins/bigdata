package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/andey-robins/bigdata/streaming/hashtable"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.Usage = func() {
		fmt.Println("Run with -help for help information")
	}

	var inFile, outFile string
	var help bool
	flag.StringVar(&inFile, "in", "./data.csv", "the file to use for data input")
	flag.StringVar(&outFile, "out", "", "the file to use for output")
	flag.BoolVar(&help, "help", false, "print help information")
	flag.Parse()

	if help {
		pad := func() {
			fmt.Printf("\n\n")
		}
		pad()
		fmt.Println(" Welcome to the streams explorer tool!")
		fmt.Print(" This code is licensed under GPLv3")
		pad()
		fmt.Println(" Args:")
		fmt.Println("  -in       The relative path to the input file. Defaults to ./data.csv")
		fmt.Println("  -out      The name of the desired output file. Omit this to print to standard output")
		fmt.Println("  -help     Display this help text :)")
		pad()
		return
	}

	driver(inFile, outFile)
}

func driver(inFile, outFile string) {
	ht := hashtable.New(int(math.Pow(2, 16)), sha256.Sum256)

	// read in and split up the data on commas and newlines
	// insert them into the hashtable
	dat, err := os.ReadFile(inFile)
	check(err)
	for _, line := range strings.Split(string(dat), "\n") {
		for _, product := range strings.Split(line, ",") {
			if product != "" {
				product = strings.TrimSpace(product)
				if !ht.Exists(product) {
					ht.Insert(product, 1)
				} else {
					prev, err := ht.Get(product)
					check(err)
					ht.Update(product, prev+1)
				}
			}
		}
	}

	if outFile == "" {
		fmt.Println(ht.Print())
	} else {
		err := os.WriteFile(outFile, []byte(ht.Print()), 0644)
		check(err)
	}
}
