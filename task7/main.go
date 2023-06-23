package main

import (
	"fmt"
	"math"
)

type ECPoint struct {
	X float64
	Y float64
}

// BasePointGGet returns the base point for the elliptic curve.
func BasePointGGet() (point ECPoint) {
	return ECPoint{X: 3, Y: 10}
}

// ECPointGen generates a new point on the elliptic curve with the given coordinates.
// @param x: X-coordinate of the point.
// @param y: Y-coordinate of the point.
// @return point: Generated ECPoint.
func ECPointGen(x, y float64) (point ECPoint) {
	return ECPoint{X: x, Y: y}
}

// IsOnCurveCheck checks if a point is on the elliptic curve.
// @param a: The ECPoint to be checked.
// @return c: True if the point is on the curve, false otherwise.
func IsOnCurveCheck(a ECPoint) (c bool) {
	// Y^2 = x^3 + X * a + b
	// where a = 0, b = 7 (EDSA bitcoin curve)
	roundValueX := math.Sqrt(math.Pow(a.X, 3) + 7)
	return roundValueX == a.Y
}

// AddECPoints adds two points on the elliptic curve.
// @param a: The first ECPoint to be added.
// @param b: The second ECPoint to be added.
// @return c: The result of the addition, a new ECPoint.
func AddECPoints(a, b ECPoint) (c ECPoint) {
	m := (a.Y - b.Y) / (a.X - b.X)
	xr := math.Pow(m, 2) - a.X - b.X
	yr := a.Y + m*(xr-a.X)
	c = ECPoint{xr, yr}
	return c
}

// DoubleECPoints doubles a point on the elliptic curve.
// @param a: The ECPoint to be doubled.
// @return c: The result of the doubling, a new ECPoint.
func DoubleECPoints(a ECPoint) (c ECPoint) {
	x2 := math.Pow(((3*math.Pow(a.X, 2)+0)/(2*a.Y)), 2) - 2*a.X
	y2 := -a.Y + ((3*math.Pow(a.X, 2)+0)/(2*a.Y))*(a.X-x2)
	c = ECPoint{x2, y2}
	return c
}

// ScalarMult multiplies a point on the elliptic curve by a scalar value.
// @param k: The scalar value to multiply the point by.
// @param a: The ECPoint to be multiplied.
// @return c: The result of the scalar multiplication, a new ECPoint.
func ScalarMult(k int, a ECPoint) (c ECPoint) {
	Xr, Yr := 0.0, 0.0

	for i := 0; i < k; i++ {
		Xr += a.X
		Yr += a.Y
	}

	return ECPointGen(Xr, Yr)
}

// PrintECPoint prints the coordinates of an ECPoint.
// @param point: The ECPoint to be printed.
func PrintECPoint(point ECPoint) {
	fmt.Printf("Point (%f, %f)\n", point.X, point.Y)
}

// IsEqual checks if two ECPoints are equal.
// @param pointFirst: The first ECPoint.
// @param pointSecond: The second ECPoint.
// @return bool: True if the points are equal, false otherwise.
func IsEqual(pointFirst ECPoint, pointSecond ECPoint) bool {
	return pointFirst.X == pointSecond.X && pointFirst.Y == pointSecond.Y
}

func main() {
	//create base point
	G := ECPoint{}
	G = BasePointGGet()
	fmt.Printf("Base point is: (%f, %f)\n", G.X, G.Y)

	P := ECPoint{}
	P = ECPointGen(1, math.Sqrt(8)) // I found this Y when I was solving on paper
	fmt.Printf("Gen point: (%f, %f)\n", P.X, P.Y)

	fmt.Println("Point P is on curve:", IsOnCurveCheck(P))

	fmt.Println("Add point", AddECPoints(P, G))

	fmt.Println("Duble point", DoubleECPoints(P))
	fmt.Println("Scalar Mult", ScalarMult(5, P))

	PrintECPoint(G)

	// G := ECPoint{}
	// G = BasePointGGet()
	// fmt.Printf("Base point is: (%f, %f)\n", G.X, G.Y)
	// d := 3
	// k := 5

	// H1 := ScalarMult(d, G)
	// H2 := ScalarMult(k, H1)

	// H3 := ScalarMult(k, G)
	// H4 := ScalarMult(d, H3)

	// fmt.Println(IsEqual(H2, H4))
}
