package main

import "github.com/shopspring/decimal"

func Arr() [1024]int {
	var x [1024]int
	for i := 0; i < len(x); i++ {
		x[i] = i
	}
	return x
}

func Slice() []int {
	x := make([]int, 1024)
	for i := 0; i < len(x); i++ {
		x[i] = i
	}
	return x
}

func Mul(f1, f2 float64) float64 {
	return f1 * f2
}

func Div(f1, f2 float64) float64 {
	return f1 / f2
}

func Add(f1, f2 float64) float64 {
	return f1 + f2
}

func Sub(f1, f2 float64) float64 {
	return f1 - f2
}

func FloatMul(f1, f2 float64) float64 {
	ff1 := decimal.NewFromFloat(f1)
	ff2 := decimal.NewFromFloat(f2)
	mul := ff1.Mul(ff2)
	r, _ := mul.Float64()
	return r
}

func FloatDiv(f1, f2 float64) float64 {
	ff1 := decimal.NewFromFloat(f1)
	ff2 := decimal.NewFromFloat(f2)
	div := ff1.Div(ff2)
	r, _ := div.Float64()
	return r
}

func FloatAdd(f1, f2 float64) float64 {
	ff1 := decimal.NewFromFloat(f1)
	ff2 := decimal.NewFromFloat(f2)
	sum := ff1.Add(ff2)
	r, _ := sum.Float64()
	return r
}

func FloatSub(f1, f2 float64) float64 {
	ff1 := decimal.NewFromFloat(f1)
	ff2 := decimal.NewFromFloat(f2)
	sum := ff1.Sub(ff2)
	r, _ := sum.Float64()
	return r
}
