package bfr

import (
	"math"
	"strings"

	"github.com/andey-robins/bigdata/clustering/data"
)

func Cluster2D(fname string) {
	dataStream := NewChunk(fname)

	discardSet := make([]*BfrClustroid, 0)
	compressedSet := make([]*BfrClustroid, 0)
	retainedSet := make([]*BfrClustroid, 0)

	// initialize 10 clustroids with random points
	// note, the points are randomized in the input
	for i := 0; i < 100; i++ {
		point := getNextPointList(dataStream)
		centroid := NewClustroid(2)
		centroid.AddPoint(point)
		discardSet = append(discardSet, centroid)
	}

	for dataStream.Continue() {
		point := getNextPointList(dataStream)
		newCluster, listCount := clusterNewPoint(discardSet, compressedSet, retainedSet, point)
		if newCluster != nil && listCount == 2 {
			retainedSet = newCluster
		} else if newCluster != nil && listCount == 1 {
			compressedSet = newCluster
		}
		mergeCompressedIntoDiscard(discardSet, compressedSet)
	}

	for _, d := range discardSet {
		d.Print()
	}
	// for _, d := range compressedSet {
	// 	fmt.Print("\n\ncompressed\n\n")
	// 	d.Print()
	// }
	// for _, d := range retainedSet {
	// 	d.Print()
	// }
}

func getNextPointList(chunk *Chunk) []int {
	// fmt.Println("next point")
	nextPointString := strings.Split(chunk.NextChunk(), ",")
	nextPoint := data.NewTwoDimData(nextPointString)
	return nextPoint.Serialize()
}

func clusterNewPoint(discards, compresses, retains []*BfrClustroid, point []int) ([]*BfrClustroid, int) {
	floatPoint := make([]float64, 0)
	for _, e := range point {
		floatPoint = append(floatPoint, float64(e))
	}

	var closestCluster *BfrClustroid = nil
	clusterIdx := -1
	closestDistance := 1_000_000_000.0
	for i, d := range discards {
		// fmt.Println(d.MahalanobisDistance(floatPoint))
		if d.MahalanobisDistance(floatPoint) <= closestDistance {
			closestDistance = d.MahalanobisDistance(floatPoint)
			closestCluster = d.Duplicate()
			clusterIdx = i
		}
	}
	if closestCluster != nil {
		closestCluster.AddPoint(point)
		// fmt.Println(closestCluster.sums)
		// sqrt to get standard deviation from variance
		if math.Sqrt(closestCluster.AverageVariance()) <= 3 {
			discards[clusterIdx] = closestCluster
			return nil, 0
		}
	}

	closestDistance = 1_000_000_000.0
	closestCluster = nil
	for i, d := range discards {
		if d.MahalanobisDistance(floatPoint) <= closestDistance {
			closestDistance = d.MahalanobisDistance(floatPoint)
			closestCluster = d.Duplicate()
			clusterIdx = i
		}
	}
	if closestCluster != nil {
		closestCluster.AddPoint(point)
		if math.Sqrt(closestCluster.AverageVariance()) <= 3 {
			discards[clusterIdx] = closestCluster
			return nil, 0
		}
	}

	for i, r := range retains {
		copy := r.Duplicate()

		if copy.AddPoint(point); math.Sqrt(copy.AverageVariance()) <= 3 {
			r.AddPoint(point)
			compresses = append(compresses, r)
			deleteAt(i, retains)
			return compresses, 1
		}
	}

	// the new point doesn't fit in any clustroid, create a new object and store it as a retains
	clustroid := NewClustroid(discards[0].dim)
	clustroid.AddPoint(point)
	retains = append(retains, clustroid)
	return retains, 2
}

func mergeCompressedIntoDiscard(discards []*BfrClustroid, compresses []*BfrClustroid) {
	mergedCompresses := make([]int, 0)
	for _, discard := range discards {
		for i, compress := range compresses {
			if discard.MahalanobisDistance(compress.Centroid()) <= 3 && !alreadyRemoved(i, mergedCompresses) {
				discard.Merge(compress)
				mergedCompresses = append(mergedCompresses, i)
			}
		}
	}

	// remove all of the merged clustroids from the compressedSet
	for i := len(mergedCompresses) - 1; i >= 0; i-- {
		deleteAt(mergedCompresses[i], compresses)
	}
}

func deleteAt(idx int, list []*BfrClustroid) []*BfrClustroid {
	if idx == len(list)-1 {
		return list[:len(list)-1]
	} else {
		return append(list[:idx], list[idx+1:]...)
	}
}

func alreadyRemoved(idx int, removedInds []int) bool {
	for _, e := range removedInds {
		if idx == e {
			return true
		}
	}
	return false
}

func Cluster10D(fname string) {
	// dataStream := NewChunk(fname)

	// discardSet := make([]BfrClustroid, 0)
	// compressedSet := make([]BfrClustroid, 0)
	// retainedSet := make([][]float64, 0)
}

func intsToFloats(ints []int) []float64 {
	floats := make([]float64, 0)
	for _, e := range ints {
		floats = append(floats, float64(e))
	}
	return floats
}
