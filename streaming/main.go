package main

import (
	"bytes"
	"crypto/sha256"
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"

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

	var seed, moments, fRange int64
	var inFile, outFile string
	var help, full, class bool
	flag.Int64Var(&seed, "seed", time.Now().Unix(), "the seed for rng")
	flag.Int64Var(&moments, "moments", 12, "the number of second moments to calculate")
	flag.Int64Var(&fRange, "range", 10_000, "the element by which all estimate moments are selected")
	flag.StringVar(&inFile, "in", "./data.csv", "the file to use for data input")
	flag.StringVar(&outFile, "out", "", "the file to use for output")
	flag.BoolVar(&full, "full", false, "run full dataset calculations")
	flag.BoolVar(&class, "class", false, "run the example from class")
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
		fmt.Println("  -seed     The seed for random number generation. Defaults to current Unix time.")
		fmt.Println("  -moments  The number of moments to use in estimating the true second moment. Defaults to 12")
		fmt.Println("  -range    The element by which all values for moments have been chose. Defaults to 10_000")
		fmt.Println("  -in       The relative path to the input file. Defaults to ./data.csv")
		fmt.Println("  -out      The name of the desired output file. Omit this to print to standard output")
		fmt.Println("  -full     Calculate the counts for the entire file instead of a simulated stream")
		fmt.Println("  -class    Run the example from class with associated defaults.")
		fmt.Println("  -help     Display this help text :)")
		pad()
		fmt.Println(" Info:")
		fmt.Println("  - The estimated moment is calculated by randomly selecting a number of indices")
		fmt.Println("      and then averaging the approximated moments for each of these indices. The")
		fmt.Println("      number of indices used in the class example was 3, but the number for this")
		fmt.Println("      homework is 12.")
		fmt.Println("  - The -full flag overrides all other flags, and since it's operations are done")
		fmt.Println("      on the complete input, other flags have no meaning when full is set.")
		pad()
		return
	}

	if full {
		fullDriver(inFile, outFile)
	} else if class {
		driver("./tests/class.csv", "", 18, 3, 14)
	} else {
		driver(inFile, outFile, seed, moments, fRange)
	}
}

func driver(inFile, outFile string, seed, momentCount, fRange int64) {
	ht := hashtable.New(int(math.Pow(2, 16)), sha256.Sum256)

	// pick 12 elements at random in the first 10_000 elements
	// for our moment estimates
	indices := make([]int, 0)
	rand.Seed(seed)
	for i := 0; i < int(momentCount); i++ {
		indices = append(indices, rand.Intn(int(fRange)))
	}
	moments := hashtable.New(3, sha256.Sum256)
	streamCardinality := 0

	// read in and split up the data on commas and newlines
	// insert them into the hashtable
	// also count the moment information to be used in the AMS algorithm
	dat, err := os.ReadFile(inFile)
	check(err)
	for _, line := range strings.Split(string(dat), "\n") {
		for _, product := range strings.Split(line, ",") {
			if product != "" {

				product = strings.TrimSpace(product)
				streamCardinality++

				// put into the complete hash table
				// this is the part that would likely be impossible for a
				// true stream
				if !ht.Exists(product) {
					ht.Insert(product, 1)
				} else {
					prev, err := ht.Get(product)
					check(err)
					ht.Update(product, prev+1)
				}

				// create entries in the moments hash table if the
				// selected element is in the indices we generated
				for _, index := range indices {
					if index == streamCardinality {
						if !moments.Exists(product) {
							moments.Insert(product, 0)
						}
					}
				}

				// update the moment hashtable if the current element is
				// in that table
				if prev, err := moments.Get(product); err == nil {
					moments.Update(product, prev+1)
				}
			}
		}
	}

	// calculate the true surprise number for the stream
	keys := ht.Keys()
	secondMoment := 0
	sCardinality := len(keys)
	for _, key := range keys {
		val, err := ht.Get(key)
		check(err)
		secondMoment += int(math.Pow(float64(val), 2))
	}

	// estimate the surprise number using the Alon-Matias-Szegedy algorithm
	keys = moments.Keys()
	estimatedMomentSum := 0
	for _, key := range keys {
		val, err := moments.Get(key)
		check(err)
		estimatedMomentSum += streamCardinality * (2*val - 1)
	}
	estimatedMoment := estimatedMomentSum / len(keys)

	// write to outfile or stdout depending on flags
	if outFile == "" {
		fmt.Printf("Stream cardinality   = %v\n", streamCardinality)
		fmt.Printf("S cardinality        = %v\n", sCardinality)
		fmt.Printf("True surprise number = %v\n", secondMoment)
		fmt.Printf("Estimated surprise   = %v\n", estimatedMoment)
	} else {
		var out bytes.Buffer

		out.WriteString(fmt.Sprintf("Stream cardinality   = %v\n", streamCardinality))
		out.WriteString(fmt.Sprintf("S cardinality        = %v\n", sCardinality))
		out.WriteString(fmt.Sprintf("True surprise number = %v\n", secondMoment))
		out.WriteString(fmt.Sprintf("Estimated surprise   = %v\n", estimatedMoment))

		err := os.WriteFile(outFile, out.Bytes(), 0644)
		check(err)
	}

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

	// write to outfile or stdout depending on flags
	if outFile == "" {
		fmt.Println(ht.Print())
	} else {
		err := os.WriteFile(outFile, []byte(ht.Print()), 0644)
		check(err)
	}
}
