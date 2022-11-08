package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/andey-robins/bigdata/similarity/hash"
	"github.com/andey-robins/bigdata/similarity/sentences"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Run with -help for help information")
	}

	var in string
	var help, logging bool
	var k, size int
	flag.StringVar(&in, "in", "sentence_files/tiny.txt", "the input file path")
	flag.IntVar(&k, "k", 0, "the distance count")
	flag.IntVar(&size, "size", 1073741824, "the number of buckets for the hashtable")
	flag.BoolVar(&help, "help", false, "use to display help text")
	flag.BoolVar(&logging, "log", false, "use to enable logging")
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
		fmt.Println("  -size     The number of buckets for the hashtable. Defaults to 2^30")
		fmt.Println("  -log      Use this flag to enable logging output")
		fmt.Println("  -help     Display this help text :)")
		pad()
		return
	}

	if !logging {
		log.SetOutput(ioutil.Discard)
	}

	if k == 0 {
		driver_0(in, size)
	} else {
		driver(in, size)
	}
}

func driver(inFile string, size int) {
	ss := sentences.New(size, hash.Campbell3)
	start := time.Now()
	ss.LoadFile(inFile)
	count := ss.CountSimilar()
	fmt.Printf("File '%s' has %v similar lines with distance 1.\n", inFile, count)
	fmt.Printf("Finished in %s\n", time.Since(start))
}

func driver_0(inFile string, size int) {
	ss := sentences.New(size, hash.Sha256Wrapper)
	start := time.Now()
	ss.LoadFile(inFile)
	count := ss.CountDupes()
	fmt.Printf("File '%s' has %v duplicate lines.\n", inFile, count)
	fmt.Printf("Finished in %s\n", time.Since(start))
}
