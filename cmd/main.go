package main

import (
	"fmt"
	square "github.com/unvise0/golang-united-school-homework-2"
)

func main() {
	fmt.Printf("CalcSquare(==4): %f\n", square.CalcSquare(0, 4))
	fmt.Printf("CalcSquare(==3): %f\n", square.CalcSquare(0, 3))
	fmt.Printf("CalcSquare(==0): %f\n", square.CalcSquare(0, 0))

	fmt.Println("-----------------------------------")

	fmt.Printf("CalcSquare(==4): %f\n", square.CalcSquare(1, 4))
	fmt.Printf("CalcSquare(==3): %f\n", square.CalcSquare(1, 3))
	fmt.Printf("CalcSquare(==0): %f\n", square.CalcSquare(1, 0))

	fmt.Println("-----------------------------------")

	fmt.Printf("CalcSquare(==4): %f\n", square.CalcSquare(3, 4))
	fmt.Printf("CalcSquare(==3): %f\n", square.CalcSquare(3, 3))
	fmt.Printf("CalcSquare(==0): %f\n", square.CalcSquare(3, 0))

	fmt.Println("-----------------------------------")

	fmt.Printf("CalcSquare(==4): %f\n", square.CalcSquare(7.23487, 4))
	fmt.Printf("CalcSquare(==3): %f\n", square.CalcSquare(7.23487, 3))
	fmt.Printf("CalcSquare(==0): %f\n", square.CalcSquare(7.23487, 0))
}
