package bfr

import (
	"errors"
	"fmt"
	"math"
)

type BfrClustroid struct {
	n      int
	dim    int
	sums   []int
	sumsSq []int
}

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

func (c *BfrClustroid) AddPoint(dims []int) {
	c.n++
	for i := 0; i < c.dim; i++ {
		c.sums[i] += dims[i]
		c.sumsSq[i] = c.sumsSq[i] + dims[i]*dims[i]
	}
	// fmt.Println(c.sums)
	// fmt.Println(c.sumsSq)
}

func (c *BfrClustroid) Centroid() []float64 {
	centroid := make([]float64, c.dim)
	for i := 0; i < c.dim; i++ {
		centroid[i] = float64(c.sums[i]) / float64(c.n)
	}
	return centroid
}

func (c *BfrClustroid) Variance() []float64 {
	variance := make([]float64, c.dim)
	for i := 0; i < c.dim; i++ {
		if c.n == 1 {
			variance[i] = 1
		} else {
			variance[i] = (float64(c.sumsSq[i]) / float64(c.n)) - math.Pow((float64(c.sums[i])/float64(c.n)), 2)
		}
	}
	return variance
}

func (c *BfrClustroid) AverageVariance() float64 {
	variance := c.Variance()
	avg := 0.0
	for _, e := range variance {
		avg += e
	}
	return avg / float64(c.dim)
}

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

func (c *BfrClustroid) MahalanobisDistance(point []float64) float64 {
	sum := 0.0
	cent := c.Centroid()
	sigma := c.Variance()
	for i, e := range point {
		if sigma[i] != 0 {
			// fmt.Printf("%v-%v / %v\n", e, cent[i], sigma[i])
			sum += math.Pow((e-cent[i])/sigma[i], 2)
		} else {
			sum += e - cent[i]
		}

	}
	// fmt.Println(sum)
	return math.Sqrt(sum)
}

func (c *BfrClustroid) Print() {
	fmt.Printf("Centroid: %v -- Std Deviation: %v \n", c.Centroid(), c.Variance())
}

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
