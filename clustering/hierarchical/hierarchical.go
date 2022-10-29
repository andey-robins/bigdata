package hierarchical

import "math"

type Cluster struct {
}

func EuclideanDistance(p1, p2 []int) float64 {
	euclideanDistSum := 0
	for i, elem := range p1 {
		diff := elem - p2[i]
		euclideanDistSum += int(math.Pow(float64(diff), float64(2)))
	}
	return math.Sqrt(float64(euclideanDistSum))
}
