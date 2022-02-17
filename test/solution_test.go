package test

import (
	square "github.com/unvise0/golang-united-school-homework-2"
	"math"
	"testing"
)

const Delta = 0.0001

type TestCalcSquareStruct struct {
	SidesLen float64
	SidesNum square.SidesNum
	Expected float64
}

func TestCalcSquare(t *testing.T) {
	var s []TestCalcSquareStruct

	t.Run("square area test group", func(t *testing.T) {
		s = []TestCalcSquareStruct{
			{SidesLen: 0, SidesNum: 4, Expected: 0},
			{SidesLen: 1, SidesNum: 4, Expected: 1},
			{SidesLen: 3, SidesNum: 4, Expected: 9},
			{SidesLen: 7.23487, SidesNum: 4, Expected: 52.3433439169},
		}
		assertCalcSquareStruct(t, s)
	})

	t.Run("equilateral triangle area test group", func(t *testing.T) {
		s = []TestCalcSquareStruct{
			{SidesLen: 0, SidesNum: 3, Expected: 0},
			{SidesLen: 1, SidesNum: 3, Expected: 0.4330127018922193},
			{SidesLen: 3, SidesNum: 3, Expected: 3.8971143170299736},
			{SidesLen: 7.23487, SidesNum: 3, Expected: 22.66533277553053},
		}
		assertCalcSquareStruct(t, s)
	})

	t.Run("circle area test group", func(t *testing.T) {
		s = []TestCalcSquareStruct{
			{SidesLen: 0, SidesNum: 0, Expected: 0},
			{SidesLen: 1, SidesNum: 0, Expected: 3.141592653589793},
			{SidesLen: 3, SidesNum: 0, Expected: 28.274333882308138},
			{SidesLen: 7.23487, SidesNum: 0, Expected: 164.441464713657},
		}
		assertCalcSquareStruct(t, s)
	})

	t.Run("Wrong number of side", func(t *testing.T) {
		s = []TestCalcSquareStruct{
			{SidesLen: 0, SidesNum: -1, Expected: 0},
			{SidesLen: -62.123, SidesNum: -123, Expected: 0},
			{SidesLen: 9221.23, SidesNum: 345, Expected: 0},
			{SidesLen: 4.8712, SidesNum: 931, Expected: 0},
		}
		assertCalcSquareStruct(t, s)
	})
}

func assertCalcSquareStruct(t *testing.T, s []TestCalcSquareStruct) {
	for _, v := range s {
		got := square.CalcSquare(v.SidesLen, v.SidesNum)
		if !assertWithDelta(v.Expected, got, Delta) {
			t.Logf("CalcSquare expected: %f, got: %f", v.Expected, got)
			t.Fail()
		}
	}
}

func assertWithDelta(a, b, delta float64) bool {
	return math.Abs(a-b) < delta
}
