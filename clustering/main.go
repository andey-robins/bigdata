package main

import (
	"flag"
	"fmt"

	"github.com/andey-robins/bigdata/clustering/bfr"
	"github.com/andey-robins/bigdata/clustering/generator"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Run with -help for help information")
	}

	var dim int
	var help, gen bool
	flag.IntVar(&dim, "dim", 2, "the dimensionality to load")
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
		fmt.Println("  -dim      The dimensionality of the data. Defaults to 2")
		fmt.Println("  -gen      Generate the datasets")
		fmt.Println("  -help     Display this help text :)")
		pad()
		return
	}

	if gen {
		generator.GenerateData2D()
		generator.GenerateDataManyDim()
	}

	if dim == 2 {
		bfr.Cluster2D("data2d.csv")
	} else if dim == 2 {
		bfr.Cluster10D("data10d.csv")
	}
}
