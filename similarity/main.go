package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime/pprof"
	"time"

	"github.com/andey-robins/bigdata/similarity/hash"
	"github.com/andey-robins/bigdata/similarity/sentences"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Run with -help for help information")
	}

	var in string
	var help, logging, experiment bool
	var k, size int
	flag.StringVar(&in, "in", "sentence_files/tiny.txt", "the input file path")
	flag.IntVar(&k, "k", 0, "the distance count")
	flag.IntVar(&size, "size", 1_000_000, "the number of buckets for the hashtable")
	flag.BoolVar(&help, "help", false, "use to display help text")
	flag.BoolVar(&logging, "log", false, "use to enable logging")
	flag.BoolVar(&experiment, "exp", false, "use to run full experiment")
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
		fmt.Println("  -size     The number of buckets for the hashtable. Defaults to 1M")
		fmt.Println("  -log      Use this flag to enable logging output")
		fmt.Println("  -exp      Use this flag to run through all files for experiment with k flag value")
		fmt.Println("  -help     Display this help text :)")
		pad()
		return
	}

	if !logging {
		log.SetOutput(ioutil.Discard)
	}

	if experiment {
		experiments := []struct {
			fname string
			size  int
		}{
			// {"sentence_files/tiny.txt", 100},
			// {"sentence_files/small.txt", 100},
			// {"sentence_files/100.txt", 100},
			// {"sentence_files/1K.txt", 1_000},
			{"sentence_files/10K.txt", 10_000},
			{"sentence_files/100K.txt", 100_000},
			// {"sentence_files/1M.txt", 1_000_000},
			// {"sentence_files/5M.txt", 5_000_000},
			// {"sentence_files/25M.txt", 25_000_000},
		}

		for _, exp := range experiments {
			cpu, err := os.Create("cpu.prof")
			if err != nil {
				log.Fatal(err)
			}
			pprof.StartCPUProfile(cpu)
			defer pprof.StopCPUProfile()

			if k == 0 {
				driver_0(exp.fname, exp.size)
			} else {
				driver(exp.fname, exp.size)
			}
		}

	} else {
		if k == 0 {
			driver_0(in, size)
		} else {
			driver(in, size)
		}
	}

}

func driver(inFile string, size int) {
	ss := sentences.New(size, hash.Campbell4)
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
