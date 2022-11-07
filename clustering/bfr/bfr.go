package bfr

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/andey-robins/bigdata/clustering/data"
)

func Cluster2D(fname string) {
	// File is an io.Reader, meaning it's a stream
	file, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	discardSet := make([]*BfrClustroid, 0)
	compressSet := make([]*BfrClustroid, 0)
	retainSet := make([]*BfrClustroid, 0)

	line := 0
	for scanner.Scan() {
		line++
		if line%10000 == 0 {
			fmt.Printf("On line %v\n", line)
		}
		// read in the next point
		text := scanner.Text()
		log.Printf("Input: %v\n", text)
		pointList := getNextPointList2D(text)
		soloClustroid := NewClustroid(len(pointList))
		soloClustroid.AddPoint(pointList)
		log.Printf("New clustroid at %v\n", soloClustroid.Centroid())

		// add the point to the retain set
		retainSet = append(retainSet, soloClustroid)

		// merge all of our sets occasionally
		if line%10_000 == 0 {
			retainSet, compressSet = mergeLowerIntoUpper(retainSet, compressSet)
			compressSet, discardSet = mergeLowerIntoUpper(compressSet, discardSet)
		}

		removedIndices := make([]int, 0)
		for i, e := range retainSet {
			if e.n > 1 {
				compressSet = append(compressSet, e)
				removedIndices = append(removedIndices, i)
			}
		}
		removeIndices(removedIndices, retainSet)
	}

	log.Println("Done with scaner")

	// ensure no errors processing our input stream
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("\n\n\n")
	fmt.Println("---------------")
	fmt.Println("| Discard set |")
	fmt.Println("---------------")
	fmt.Printf("\n\n\n")

	// output
	outputVals := make([]*BfrClustroid, 0)
	for _, e := range discardSet {
		alreadyOutput := false
		for _, out := range outputVals {
			if e == out {
				alreadyOutput = true
			}
		}
		if !alreadyOutput {
			e.Print()
		}
		outputVals = append(outputVals, e)
	}

	fmt.Printf("\n\n\n")
	fmt.Println("----------------")
	fmt.Println("| Compress set |")
	fmt.Println("----------------")
	fmt.Printf("\n\n\n")

	outputVals = make([]*BfrClustroid, 0)
	for _, e := range compressSet {
		alreadyOutput := false
		for _, out := range outputVals {
			if e == out {
				alreadyOutput = true
			}
		}
		if !alreadyOutput {
			e.Print()
		}
		outputVals = append(outputVals, e)
	}

	fmt.Printf("\n\n\n")
	fmt.Println("--------------")
	fmt.Println("| Retain set |")
	fmt.Println("--------------")
	fmt.Printf("\n\n\n")

	outputVals = make([]*BfrClustroid, 0)
	for _, e := range retainSet {
		alreadyOutput := false
		for _, out := range outputVals {
			if e == out {
				alreadyOutput = true
			}
		}
		if !alreadyOutput {
			e.Print()
		}
		outputVals = append(outputVals, e)
	}
}

func getNextPointList2D(s string) []int {
	nextPointString := strings.Split(s, ",")
	nextPoint := data.NewTwoDimData(nextPointString)
	return nextPoint.Serialize()
}

func getNextPointList10D(s string) []int {
	nextPointString := strings.Split(s, ",")
	nextPoint := data.NewTenDimData(nextPointString)
	return nextPoint.Serialize()
}

func mergePointIfClose(point *BfrClustroid, clustroids []*BfrClustroid) ([]*BfrClustroid, bool) {
	// for each clustroid, if point has a malonois under a threshhold, merge it to the clustroid
	for _, clustroid := range clustroids {
		if clustroid.MahalanobisDistance(point.Centroid()) <= 4 {
			clustroid.Merge(point)
			log.Printf("%v merged into %v dist %v\n", point.Centroid(), clustroid.Centroid(), clustroid.MahalanobisDistance(point.Centroid()))
			return clustroids, true
		}
	}
	return clustroids, false
}

func mergeLowerIntoUpper(lower []*BfrClustroid, upper []*BfrClustroid) ([]*BfrClustroid, []*BfrClustroid) {
	// for each in retain, if it has a malonois under a threshhold, merge it into the compress
	removedIndices := make([]int, 0)
	for i, retainClustroid := range lower {
		if set, merged := mergePointIfClose(retainClustroid, upper); merged {
			upper = set
			removedIndices = append(removedIndices, i)
		}
	}

	// remove the clustroids in th retain set at all of the indices in removedIndices
	lower = removeIndices(removedIndices, lower)
	lower = combineCloseClustroids(lower)
	upper = combineCloseClustroids(upper)
	return lower, upper
}

func combineCloseClustroids(clustroids []*BfrClustroid) []*BfrClustroid {
	// if two clustroids overlap, merge them
	removedIndices := make([]int, 0)
	for i, clustroidOne := range clustroids {
		for j := i + 1; j < len(clustroids); j++ {
			clustroidTwo := clustroids[j]
			if clustroidOne.MahalanobisDistance(clustroidTwo.Centroid()) <= 4 {
				clustroidTwo.Merge(clustroidOne)
				log.Printf("merged %v into %v dist %v\n", clustroidOne.Centroid(), clustroidTwo.Centroid(), clustroidOne.MahalanobisDistance(clustroidTwo.Centroid()))
				removedIndices = append(removedIndices, i)
			} else {
				log.Printf("merged %v into %v dist %v\n", clustroidOne.Centroid(), clustroidTwo.Centroid(), clustroidOne.MahalanobisDistance(clustroidTwo.Centroid()))
			}
		}
	}
	removeIndices(removedIndices, clustroids)
	return clustroids
}

func removeIndices(indices []int, list []*BfrClustroid) []*BfrClustroid {
	if len(indices) == 0 {
		return list
	}
	lastIdx := indices[len(indices)-1]
	newIndices := indices[:len(indices)-1]
	if lastIdx == len(list) {
		return removeIndices(newIndices, list[:lastIdx])
	}
	return removeIndices(newIndices, append(list[:lastIdx], list[lastIdx+1:]...))
}

func Cluster10D(fname string) {
	// File is an io.Reader, meaning it's a stream
	file, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	discardSet := make([]*BfrClustroid, 0)
	compressSet := make([]*BfrClustroid, 0)
	retainSet := make([]*BfrClustroid, 0)

	line := 0
	for scanner.Scan() {
		line++
		if line%1000 == 0 {
			fmt.Println(line)
		}
		// read in the next point
		log.Printf("Input: %v\n", scanner.Text())
		pointList := getNextPointList10D(scanner.Text())
		soloClustroid := NewClustroid(len(pointList))
		soloClustroid.AddPoint(pointList)
		log.Printf("New clustroid at %v\n", soloClustroid.Centroid())

		// add the point to the retain set
		retainSet = append(retainSet, soloClustroid)

		if line%1_000 == 0 {
			retainSet, compressSet = mergeLowerIntoUpper(retainSet, compressSet)
			compressSet, discardSet = mergeLowerIntoUpper(compressSet, discardSet)
		}

		// merge all of our sets
		retainSet, compressSet = mergeLowerIntoUpper(retainSet, compressSet)
		compressSet, discardSet = mergeLowerIntoUpper(compressSet, discardSet)

		removedIndices := make([]int, 0)
		for i, e := range retainSet {
			if e.n > 1 {
				compressSet = append(compressSet, e)
				removedIndices = append(removedIndices, i)
			}
		}
		removeIndices(removedIndices, retainSet)
	}

	log.Println("Done with scaner")

	// ensure no errors processing our input stream
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// fmt.Printf("\n\n\n")
	// fmt.Println("---------------")
	// fmt.Println("| Discard set |")
	// fmt.Println("---------------")
	// fmt.Printf("\n\n\n")

	// // output
	// outputVals := make([]*BfrClustroid, 0)
	// for _, e := range discardSet {
	// 	alreadyOutput := false
	// 	for _, out := range outputVals {
	// 		if e == out {
	// 			alreadyOutput = true
	// 		}
	// 	}
	// 	if !alreadyOutput {
	// 		e.Print()
	// 	}
	// 	outputVals = append(outputVals, e)
	// }

	fmt.Printf("\n\n\n")
	fmt.Println("----------------")
	fmt.Println("| Compress set |")
	fmt.Println("----------------")
	fmt.Printf("\n\n\n")

	outputVals := make([]*BfrClustroid, 0)
	for _, e := range compressSet {
		alreadyOutput := false
		for _, out := range outputVals {
			if e == out {
				alreadyOutput = true
			}
		}
		if !alreadyOutput {
			e.Print()
		}
		outputVals = append(outputVals, e)
	}

	// fmt.Printf("\n\n\n")
	// fmt.Println("--------------")
	// fmt.Println("| Retain set |")
	// fmt.Println("--------------")
	// fmt.Printf("\n\n\n")

	// outputVals = make([]*BfrClustroid, 0)
	// for _, e := range retainSet {
	// 	alreadyOutput := false
	// 	for _, out := range outputVals {
	// 		if e == out {
	// 			alreadyOutput = true
	// 		}
	// 	}
	// 	if !alreadyOutput {
	// 		e.Print()
	// 	}
	// 	outputVals = append(outputVals, e)
	// }
}
