package main

import (
	"fmt"
	"math"
)

func main() {
	var areaCircle float64
	fmt.Print("Please enter area of the circle ")
	fmt.Scanf("%f\n", &areaCircle)
	fmt.Printf("Diameter of the circle is equal %.3f\n", math.Sqrt(4*areaCircle/math.Pi))
	fmt.Printf("Circumference of the circle is equal %.3f\n", math.Sqrt(4*math.Pi*areaCircle))
}
