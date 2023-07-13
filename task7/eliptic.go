package eleptic

import "math/big"

type ECPoint struct {
	X *big.Int
	Y *big.Int
}

func BasePointGGet() (point ECPoint) {
	point = ECPoint{
		X: big.NewInt(11570),
		Y: big.NewInt(42257),
	}
	return point
}

func ECPointGen(x, y *big.Int) (point ECPoint) {
	point = ECPoint{
		X: x,
		Y: y,
	}
	return point
}

func IsOnCurveCheck(a ECPoint) bool {
	if a.X == nil || a.Y == nil {
		return false
	}

	left := new(big.Int).Exp(a.Y, big.NewInt(2), nil)

	right := new(big.Int).Exp(a.X, big.NewInt(3), nil)
	right.Add(right, big.NewInt(1))

	return left.Cmp(right) == 0
}

func AddECPoints(a, b ECPoint) (c ECPoint) {
	sumX := new(big.Int)
	sumX.Add(a.X, b.X)

	sumY := new(big.Int)
	sumY.Add(a.Y, b.Y)

	return ECPoint{
		X: sumX,
		Y: sumY,
	}
}

func DoubleECPoints(a ECPoint) (c ECPoint) {
	sumX := new(big.Int)
	sumX.Add(a.X, a.X)

	sumY := new(big.Int)
	sumY.Add(a.Y, a.Y)

	return ECPoint{
		X: sumX,
		Y: sumY,
	}
}
func IsEqual(a, b ECPoint) bool {
	return a.X.Cmp(b.X) == 0 && a.Y.Cmp(b.Y) == 0
}

func ScalarMult(k big.Int, a ECPoint) ECPoint {
	if a.X == nil || a.Y == nil {
		return ECPoint{}
	}

	result := ECPoint{
		X: new(big.Int),
		Y: new(big.Int),
	}

	result.X.SetInt64(0)
	result.Y.SetInt64(0)

	binaryK := k.Text(2)

	for i := 0; i < len(binaryK); i++ {
		result = AddECPoints(result, result)

		if binaryK[i] == '1' {
			result = AddECPoints(result, a)
		}
	}

	return result
}

func MultiplyECPoint(k big.Int, a ECPoint) ECPoint {
	result := ECPoint{}

	if a.X == nil || a.Y == nil {
		return result
	}

	currentPoint := a

	for i := k.BitLen() - 1; i >= 0; i-- {
		currentPoint = DoubleECPoints(currentPoint)

		if k.Bit(i) == 1 {
			currentPoint = AddECPoints(currentPoint, a)
		}
	}

	result.X = new(big.Int).Set(currentPoint.X)
	result.Y = new(big.Int).Set(currentPoint.Y)

	return result
}
