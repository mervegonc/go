package main

import (
	"fmt"
	"math"
)

func main() {
	var side1, side2, hypotenuse float64

	fmt.Println("welcome to right triangle hypotenuse calculator")
	fmt.Print("enter the first side:")
	fmt.Scan(&side1)
	fmt.Print("enter the second side")
	fmt.Scan(&side2)

	hypotenuse = math.Sqrt(math.Pow(side1, 2) + math.Pow(side2, 2))
	fmt.Printf("\nHypotenuse: %f", hypotenuse)

}
