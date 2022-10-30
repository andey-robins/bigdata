package data

import "strconv"

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
