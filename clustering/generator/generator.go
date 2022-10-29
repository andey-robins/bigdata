package generator

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/andey-robins/bigdata/clustering/data"
	"github.com/gocarina/gocsv"
)

type TwoDimData = data.TwoDimData
type TenDimData = data.TenDimData

// GenerateData2D generates a large 2D dataset with built in clusters
// and random noise
func GenerateData2D() {
	if _, err := os.Stat("data2d.csv"); err == nil {
		fmt.Println("data2d.csv already exists. Generation is deterministic, so not regenerating.\n\nIf you want to regenerate the file, please delete it first.")
		return
	}

	// setup deterministic data generation
	rand.Seed(2)

	data := make([]*TwoDimData, 0)

	// generate 10 clusters (25% of the data)
	for clusters := 0; clusters < 10; clusters++ {
		clusterX := rand.Intn(10_000 - 50)
		clusterY := rand.Intn(10_000 - 50)

		for x := clusterX; x < 50+clusterX; x++ {
			for y := clusterY; y < 50+clusterY; y++ {
				data = append(data, &TwoDimData{X: x, Y: y})
			}
		}
	}

	// generate a bunch of noise (75% of the data)
	for i := 0; i < 75_000; i++ {
		data = append(data, &TwoDimData{X: rand.Intn(10_000), Y: rand.Intn(10_000)})
	}

	// shuffle the data so clusters aren't near each other in the file
	rand.Shuffle(len(data), func(i, j int) { data[i], data[j] = data[j], data[i] })

	// write the data out to a file
	dataFile, err := os.OpenFile("data2d.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer dataFile.Close()
	err = gocsv.MarshalFile(&data, dataFile)
	if err != nil {
		panic(err)
	}
}

// GenerateDataManyDim generates a high dimensional dataset
// with 10 dimensions (10 was chosen arbitrarily)
func GenerateDataManyDim() {
	if _, err := os.Stat("data10d.csv"); err == nil {
		fmt.Println("data10d.csv already exists. Generation is deterministic, so not regenerating.\n\nIf you want to regenerate the file, please delete it first.")
		return
	}

	// setup deterministic data generation
	rand.Seed(2)

	data := make([]*TenDimData, 0)

	// generate 10 clusters (20% of the data)
	for clusters := 0; clusters < 10; clusters++ {

		centroid := make([]int, 10)
		for i := 0; i < 10; i++ {
			centroid[i] = rand.Intn(1000 - 10)
		}

		// change a variable number of the points in our cluster
		for dimensionsToChange := 1; dimensionsToChange <= 10; dimensionsToChange++ {
			// generate a number of these modified points
			for iter := 0; iter < 100; iter++ {
				// reset our clusterPoint to centroid
				clusterPoint := copyPoint(centroid)

				// identify a subset of indices to change
				indicesToChange := make([]int, dimensionsToChange)
				for i := range indicesToChange {
					indicesToChange[i] = rand.Intn(10)
				}

				// only change by a value between 1 and 5
				changeBy := rand.Intn(5)

				// apply the change
				for _, index := range indicesToChange {
					clusterPoint[index] += changeBy
				}

				// save the modified point to the struct being written to csv
				copy := copyPoint(clusterPoint)
				dataPoint := &TenDimData{
					A: copy[0],
					B: copy[1],
					C: copy[2],
					D: copy[3],
					E: copy[4],
					F: copy[5],
					G: copy[6],
					H: copy[7],
					I: copy[8],
					J: copy[9],
				}
				data = append(data, dataPoint)
			}
		}

	}

	// generate a bunch of noise (80% of the data)
	for i := 0; i < 40_000; i++ {
		data = append(data, &TenDimData{
			A: rand.Intn(1000),
			B: rand.Intn(1000),
			C: rand.Intn(1000),
			D: rand.Intn(1000),
			E: rand.Intn(1000),
			F: rand.Intn(1000),
			G: rand.Intn(1000),
			H: rand.Intn(1000),
			I: rand.Intn(1000),
			J: rand.Intn(1000),
		})
	}

	// shuffle the data so clusters aren't near each other in the file
	rand.Shuffle(len(data), func(i, j int) { data[i], data[j] = data[j], data[i] })

	// write the data out to a file
	dataFile, err := os.OpenFile("data10d.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer dataFile.Close()
	err = gocsv.MarshalFile(&data, dataFile)
	if err != nil {
		panic(err)
	}
}

func copyPoint(base []int) []int {
	work := make([]int, 10)
	for i := 0; i < 10; i++ {
		work[i] = base[i]
	}

	return work
}
