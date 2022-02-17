package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type SidesNum int

const (
	SidesCircle   SidesNum = 0
	SidesTriangle SidesNum = 3
	SidesSquare   SidesNum = 4
)

func CalcSquare(sideLen float64, sidesNum SidesNum) float64 {
	if int(sidesNum) == 4 {
		return math.Pow(sideLen, 2)
	} else if int(sidesNum) == 3 {
		return math.Pow(sideLen, 2) * math.Sqrt(3) / 4
	} else if int(sidesNum) == 0 {
		return math.Pi * math.Pow(sideLen, 2)
	}
	return 0
}
