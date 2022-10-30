package bfr

import (
	"bufio"
	"fmt"
	"os"
)

type Chunk struct {
	fname   string
	line    int
	maxLine int
	lines   []string
}

func NewChunk(fname string) *Chunk {
	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	maxLine := 0
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		maxLine++
	}

	line := 0
	return &Chunk{
		fname,
		line,
		maxLine,
		lines,
	}
}

func (c *Chunk) NextChunk() string {
	c.line++
	return c.lines[c.line]
}

func (c *Chunk) Continue() bool {
	if c.line%1_000 == 0 {
		fmt.Printf("On line %v\n", c.line)
	}
	return c.line < c.maxLine-1
}
