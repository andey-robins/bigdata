package main

import (
	"flag"
	"fmt"

	"github.com/andey-robins/bigdata/clustering/generator"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Run with -help for help information")
	}

	var help, gen bool
	flag.BoolVar(&help, "help", false, "print help information")
	flag.BoolVar(&gen, "gen", false, "generate datasets")
	flag.Parse()

	if help {
		pad := func() {
			fmt.Printf("\n\n")
		}
		pad()
		fmt.Println(" Welcome to the parallel clustering explorer tool!")
		fmt.Print(" This code is licensed under GPLv3")
		pad()
		fmt.Println(" Args:")
		fmt.Println("  -gen      Generate the datasets")
		fmt.Println("  -help     Display this help text :)")
		pad()
		return
	}

	if gen {
		generator.GenerateData2D()
		generator.GenerateDataManyDim()
	}
}
