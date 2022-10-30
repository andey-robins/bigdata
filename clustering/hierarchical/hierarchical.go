package hierarchical

import "math"

type Cluster struct {
}

func EuclideanDistance(p1, p2 []float64) float64 {
	euclideanDistSum := 0.0
	for i, elem := range p1 {
		diff := elem - p2[i]
		euclideanDistSum += math.Pow(diff, float64(2))
	}
	return math.Sqrt(euclideanDistSum)
}
