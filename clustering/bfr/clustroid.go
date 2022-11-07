package bfr

import (
	"errors"
	"fmt"
	"math"
)

// BfrClustroid is a structure used to represent a BfrClustroid
// for use with the BFR clustering algorithm
type BfrClustroid struct {
	n      int
	dim    int
	sums   []int
	sumsSq []int
}

// NewClustroid creates a new clustroid object with dimensionality dim
func NewClustroid(dim int) *BfrClustroid {
	sums := make([]int, dim)
	sumsSq := make([]int, dim)
	return &BfrClustroid{
		n:      0,
		dim:    dim,
		sums:   sums,
		sumsSq: sumsSq,
	}
}

// AddPoint puts a new point represented by the list dims into a clustroid
func (c *BfrClustroid) AddPoint(dims []int) {
	c.n++
	for i := 0; i < c.dim; i++ {
		c.sums[i] += dims[i]
		c.sumsSq[i] = c.sumsSq[i] + dims[i]*dims[i]
	}
}

// Centroid returns the point of the centroid for the clustroid
func (c *BfrClustroid) Centroid() []float64 {
	centroid := make([]float64, c.dim)
	for i := 0; i < c.dim; i++ {
		centroid[i] = float64(c.sums[i]) / float64(c.n)
	}
	return centroid
}

// Variance returns the variance for the clustroid in each dimension
func (c *BfrClustroid) Variance() []float64 {
	variance := make([]float64, c.dim)
	for i := 0; i < c.dim; i++ {
		variance[i] = (float64(c.sumsSq[i]) / float64(c.n)) - math.Pow((float64(c.sums[i])/float64(c.n)), 2)
		if variance[i] == 0 {
			variance[i] = 1
		}
	}
	return variance
}

// StandardDeviation returns the standard deviation for the clustroid
// in each dimension
func (c *BfrClustroid) StandardDeviation() []float64 {
	variance := c.Variance()
	stdDev := make([]float64, c.dim)
	for i := 0; i < c.dim; i++ {
		stdDev[i] = math.Sqrt(variance[i])
	}
	return stdDev
}

// Merge takes the parameter clustroid and merges it with the object clustroid
// returns an error if you attempt to merge two differently sized clustroids
func (c *BfrClustroid) Merge(c2 *BfrClustroid) error {
	if c.dim != c2.dim {
		return errors.New("clustroids have different dimensions")
	}

	// merge the datastructures
	c.n += c2.n
	for i, e := range c2.sums {
		c.sums[i] += e
	}
	for i, e := range c2.sumsSq {
		c.sumsSq[i] += e
	}

	return nil
}

// MahalanobisDistance calculates the Mdist between a point (argument)
// and the cluster (object). The distance is returned as a float representing
// the number of standard deviations the point is from the centroid
func (c *BfrClustroid) MahalanobisDistance(point []float64) float64 {
	sum := 0.0
	cent := c.Centroid()
	variance := c.Variance()
	for i, e := range point {
		sum += math.Pow(((e - cent[i]) / math.Sqrt(variance[i])), 2)
	}
	return math.Sqrt(sum)
}

// Print is a nice printer which outputs the centroid and standard deviation for a clustroid
func (c *BfrClustroid) Print() {
	fmt.Printf("Centroid: %v -- Variance: %v \n", c.Centroid(), c.Variance())
}

// Duplicate creates a copy of a clustroid object
func (c *BfrClustroid) Duplicate() *BfrClustroid {
	dupeSums := make([]int, c.dim)
	dupeSumsSq := make([]int, c.dim)

	for i := 0; i < c.dim; i++ {
		dupeSums[i] = c.sums[i]
		dupeSumsSq[i] = c.sumsSq[i]
	}

	return &BfrClustroid{
		n:      c.n,
		dim:    c.dim,
		sums:   dupeSums,
		sumsSq: dupeSumsSq,
	}
}
