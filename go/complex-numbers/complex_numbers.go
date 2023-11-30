package complexnumbers

import (
	"math"
)

type Number struct {
	R float64
	I float64
}

func (n Number) Real() float64 {
	return n.R
}

func (n Number) Imaginary() float64 {
	return float64(n.I)
}

func (n1 Number) Add(n2 Number) Number {
	var result Number
	result.R = n1.R + n2.R
	result.I = n1.I + n2.I

	return result
}

func (n1 Number) Subtract(n2 Number) Number {
	var result Number
	result.R = n1.R - n2.R
	result.I = n1.I - n2.I

	return result
}

func (n1 Number) Multiply(n2 Number) Number {
	var result Number
	real:=n1.R*n2.R + (n1.I * n2.I *-1)
	img:=n1.R*n2.I + n1.I*n2.R
	result.R = real
	result.I = img

	return result
}

func (n Number) Times(factor float64) Number {
	return n.Multiply(Number{R: factor, I: 0})
}

func (n1 Number) Divide(n2 Number) Number {
	var result Number
	n2C:=n2.Conjugate()
	divisor:=n2.Multiply(n2C)
	n1N:=n1.Multiply(n2C)
	result.R = n1N.R / divisor.R
	result.I = n1N.I / divisor.R

	return result
}

func (n Number) Conjugate() Number {
	result:=Number{R: n.R, I: n.I *-1}
	return result
}

func (n Number) Abs() float64 {
	return math.Sqrt(math.Pow(n.R, 2) + math.Pow(n.I, 2))
}

func Rect2Polar(n Number) Number {
	z:= math.Sqrt(math.Pow(n.R, 2) + math.Pow(n.I, 2))
	argument:=math.Atan(n.I/n.R)

	return Number{z, argument}
}

func Polar2Rect(n Number) Number {
	x:=math.Cos( n.I ) 
	y:=math.Sin( n.I )

	return Number{n.R, x + y}
}

func RadToGrad(value float64) float64 {
	return (value / math.Pi) * 180
}

func (n Number) Exp() Number {
	return Number{math.Exp(n.R) * math.Cos(n.I), math.Exp(n.R) * math.Sin(n.I)}
}
