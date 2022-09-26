package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/chrislusf/glow/flow"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Run with -help for help information")
	}

	var help, wc bool
	flag.BoolVar(&wc, "wc", false, "use to run wc test")
	flag.BoolVar(&help, "help", false, "use to display help text")
	flag.Parse()

	if help {
		pad := func() {
			fmt.Printf("\n\n")
		}
		pad()
		fmt.Println(" Welcome to the mapreduce explorer tool!")
		fmt.Println(" This code is build on the Glow library")
		fmt.Print(" This code is licensed under GPLv3")
		pad()
		fmt.Println(" Args:")
		fmt.Println("  -wc       Run the word count map reduce program on julius-caesar.txt")
		fmt.Println("  -help     Display this help text :)")
		pad()
		return
	}

	if wc {
		wordCount()
	}

}

func wordCount() {
	flow.New().TextFile(
		"julius-caesar.txt", 4,
	).Map(func(line string, ch chan string) {
		for _, word := range strings.Split(line, " ") {
			if word != "" {
				ch <- word
			}
		}
	}).Map(func(key string) int {
		return 1
	}).Reduce(func(x, y int) int {
		return x + y
	}).Map(func(x int) {
		fmt.Printf("Words = %v\n", x)
	}).Run()
}
