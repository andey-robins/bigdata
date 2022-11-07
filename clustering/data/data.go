package data

import "strconv"

type Dimensional interface {
	TwoDimData
	TenDimData
}

type TwoDimData struct {
	X int `csv:"x"`
	Y int `csv:"y"`
}

func NewTwoDimData(vals []string) *TwoDimData {
	x, _ := strconv.Atoi(vals[0])
	y, _ := strconv.Atoi(vals[1])

	return &TwoDimData{
		X: x,
		Y: y,
	}
}

func (d *TwoDimData) Serialize() []int {
	data := make([]int, 2)
	data[0] = d.X
	data[1] = d.Y
	return data
}

type TenDimData struct {
	A int `csv:"a"`
	B int `csv:"b"`
	C int `csv:"c"`
	D int `csv:"d"`
	E int `csv:"e"`
	F int `csv:"f"`
	G int `csv:"g"`
	H int `csv:"h"`
	I int `csv:"i"`
	J int `csv:"j"`
}

func NewTenDimData(vals []string) *TenDimData {
	a, _ := strconv.Atoi(vals[0])
	b, _ := strconv.Atoi(vals[1])
	c, _ := strconv.Atoi(vals[2])
	d, _ := strconv.Atoi(vals[3])
	e, _ := strconv.Atoi(vals[4])
	f, _ := strconv.Atoi(vals[5])
	g, _ := strconv.Atoi(vals[6])
	h, _ := strconv.Atoi(vals[7])
	i, _ := strconv.Atoi(vals[8])
	j, _ := strconv.Atoi(vals[9])

	return &TenDimData{
		A: a,
		B: b,
		C: c,
		D: d,
		E: e,
		F: f,
		G: g,
		H: h,
		I: i,
		J: j,
	}
}

func (d *TenDimData) Serialize() []int {
	data := make([]int, 10)
	data[0] = d.A
	data[1] = d.B
	data[2] = d.C
	data[3] = d.D
	data[4] = d.E
	data[5] = d.F
	data[6] = d.G
	data[7] = d.H
	data[8] = d.I
	data[9] = d.J
	return data
}
