package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Run with -help for help information")
	}

	var help bool
	flag.BoolVar(&help, "help", false, "print help information")

	if help {
		pad := func() {
			fmt.Printf("\n\n")
		}
		pad()
		fmt.Println(" Welcome to the streams explorer tool!")
		fmt.Print(" This code is licensed under GPLv3")
		pad()
		fmt.Println(" Args:")
		fmt.Println("  -help     Display this help text :)")
		pad()
		return
	}

	driver()
}

func driver() {

}
