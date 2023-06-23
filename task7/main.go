package main

import (
	"fmt"
	"math"
)

type ECPoint struct {
	X float64
	Y float64
}

func BasePointGGet() (point ECPoint) {
	return ECPoint{X: 3, Y: 10}
} //G-generator receiving

func ECPointGen(x, y float64) (point ECPoint) {
	return ECPoint{X: x, Y: y}
}

func IsOnCurveCheck(a ECPoint) (c bool) {
	//Y^2 = x^3 + X * a + b
	// where a = 0, b = 7 (EDSA bitcoin curve)
	roundValueX := math.Sqrt(math.Pow(a.X, 3) + 7)
	return roundValueX == a.Y

}

func AddECPoints(a, b ECPoint) (c ECPoint) {
	m := (a.Y - b.Y) / (a.X - b.X)
	xr := math.Pow(m, 2) - a.X - b.X
	yr := a.Y + m*(xr-a.X)
	c = ECPoint{xr, yr}
	return c
} //P + Q

func DoubleECPoints(a ECPoint) (c ECPoint) {
	x2 := math.Pow(((3*math.Pow(a.X, 2)+0)/(2*a.Y)), 2) - 2*a.X
	y2 := -a.Y + ((3*math.Pow(a.X, 2)+0)/(2*a.Y))*(a.X-x2)
	c = ECPoint{x2, y2}
	return c
}
func ScalarMult(k int, a ECPoint) (c ECPoint) {
	Xr, Yr := 0.0, 0.0 

	for i := 0; i < k; i++ {
		Xr += a.X
		Yr += a.Y
	}

	return ECPointGen(Xr, Yr)
}

func PrintECPoint(point ECPoint) {
	fmt.Printf("Point (%f, %f)\n", point.X, point.Y)
}

func IsEqual(pointFirst ECPoint, pointSecond ECPoint) bool {
	return pointFirst.X == pointSecond.X && pointFirst.Y == pointSecond.Y
}

func main() {
	// create base point
	// G := ECPoint{}
	// G = BasePointGGet()
	// fmt.Printf("Base point is: (%f, %f)\n", G.X, G.Y)

	// P := ECPoint{}
	// P = ECPointGen(1, math.Sqrt(8)) // I found this Y when I was solving on paper
	// fmt.Printf("Gen point: (%f, %f)\n", P.X, P.Y)

	// fmt.Println("Point P is on curve:", IsOnCurveCheck(P))

	// fmt.Println("Add point", AddECPoints(P, G))

	// fmt.Println("Duble point", DoubleECPoints(P))
	// fmt.Println("Scalar Mult", ScalarMult(5, P))

	// PrintECPoint(G)

	G := ECPoint{}
	G = BasePointGGet()
	fmt.Printf("Base point is: (%f, %f)\n", G.X, G.Y)
	d := 3
	k := 5

	H1 := ScalarMult(d, G)
	H2 := ScalarMult(k, H1)

	H3 := ScalarMult(k, G)
	H4 := ScalarMult(d, H3)

	fmt.Println(IsEqual(H2, H4))

}
