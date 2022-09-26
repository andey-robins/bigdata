package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/chrislusf/glow/flow"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Run with -help for help information")
	}

	var in, out string
	var help, wc, rbow bool
	flag.StringVar(&in, "in", "pass/common.txt", "the input file path")
	flag.StringVar(&out, "out", "rainbow.csv", "the output file path")
	flag.BoolVar(&rbow, "rbow", false, "use to generate the rainbow table")
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
		fmt.Println("  -rbow     Generate the rainbow table for the password list")
		fmt.Println("  -in       The file name for password input. Defaults to pass/common.txt")
		fmt.Println("  -out      The file name for the rainbow table. Defaults to rainbow.txt")
		fmt.Println("  -wc       Run the word count map reduce program on julius-caesar.txt")
		fmt.Println("  -help     Display this help text :)")
		pad()
		return
	}

	if wc {
		wordCount()
	} else if rbow {
		rainbow(in, out)
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

func rainbow(in, out string) {
	flow.New().TextFile(
		in, 10,
	).Map(func(line string, ch chan string) {
		ch <- line
	}).Map(func(key string) string {
		md5Hash := md5.Sum([]byte(key))
		sha1Hash := sha1.Sum([]byte(key))
		sha2Hash := sha256.Sum256([]byte(key))
		sha3Hash := sha512.Sum512_256([]byte(key))
		return fmt.Sprintf("%v,%v,%v,%v,%v\n", key, hex.EncodeToString(md5Hash[:]), hex.EncodeToString(sha1Hash[:]), hex.EncodeToString(sha2Hash[:]), hex.EncodeToString(sha3Hash[:]))
	}).Reduce(func(h1, h2 string) string {
		return fmt.Sprintf("%v%v", h1, h2)
	}).Map(func(res string) {
		err := os.WriteFile(out, []byte(res), 0644)
		if err != nil {
			panic(err)
		}
	}).Run()
}
