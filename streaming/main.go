package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"math"
	"math/rand"
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
	var help, full bool
	flag.StringVar(&inFile, "in", "./data.csv", "the file to use for data input")
	flag.StringVar(&outFile, "out", "", "the file to use for output")
	flag.BoolVar(&full, "full", false, "run full dataset calculations")
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
		fmt.Println("  -full     Calculate the results for the entire file instead of a simulated stream")
		fmt.Println("  -help     Display this help text :)")
		pad()
		return
	}

	if full {
		fullDriver(inFile, outFile)
	} else {
		driver(inFile, outFile)
	}
}

func driver(inFile, outFile string) {
	ht := hashtable.New(int(math.Pow(2, 16)), sha256.Sum256)

	// start := 0
	// end := 10
	start := rand.Intn(50_000)
	end := rand.Intn(50_000) + 50_000

	// read in and split up the data on commas and newlines
	// insert them into the hashtable
	dat, err := os.ReadFile(inFile)
	check(err)
	for i, line := range strings.Split(string(dat), "\n") {
		if i >= start && i <= end {
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
	}

	keys := ht.Keys()

	streamCardinality := 0
	secondMoment := 0
	sCardinality := len(keys)
	for _, key := range keys {
		val, err := ht.Get(key)
		check(err)
		streamCardinality += val
		secondMoment += int(math.Pow(float64(val), 2))
	}

	fmt.Printf("Stream card      = %v\n", streamCardinality)
	fmt.Printf("S card           = %v\n", sCardinality)
	fmt.Printf("True moment      = %v\n", secondMoment)
	fmt.Printf("Estimated moment = %v\n", 0)
}

func fullDriver(inFile, outFile string) {
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
