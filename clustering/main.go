package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/andey-robins/bigdata/clustering/bfr"
	"github.com/andey-robins/bigdata/clustering/generator"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Run with -help for help information")
	}

	var dim int
	var help, gen, logging bool
	flag.IntVar(&dim, "dim", 2, "the dimensionality to load")
	flag.BoolVar(&help, "help", false, "print help information")
	flag.BoolVar(&gen, "gen", false, "generate datasets")
	flag.BoolVar(&logging, "log", false, "enable logging")
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

	if !logging {
		log.SetOutput(ioutil.Discard)
	}

	if gen {
		generator.GenerateData2D()
		generator.GenerateDataManyDim()
	}

	// files := []string{
	// 	"data2d.csv",
	// 	"cheung_2d.csv",
	// 	"dowalter-2d.csv",
	// 	"jarhan_2d.csv",
	// 	"joshi_2d-1.csv",
	// 	"kalivoda_2d.csv",
	// 	"liu_2d.csv",
	// 	"mcilwaine_2d.csv",
	// 	"shah_2d.csv",
	// 	"tan_2d.csv",
	// 	"wu_2d.csv",
	// }

	// for _, s := range files {
	// 	fmt.Printf("%v\n", s)
	// 	bfr.Cluster2D(s)
	// }

	files := []string{
		"data10d.csv",
		"dowalter_10d.csv",
		"cheung10d.csv",
		"mcilwaine_10d.csv",
		"tan_10d.csv",
		"jarhan_10d.csv",
	}

	for _, s := range files {
		fmt.Printf("%v\n", s)
		bfr.Cluster10D(s)
	}

	// if dim == 2 {
	// 	bfr.Cluster2D("data2d.csv")
	// } else if dim == 10 {
	// 	bfr.Cluster10D("data10d.csv")
	// } else if dim == 1 {
	// 	bfr.Cluster2D("data1d.csv")
	// }
}
