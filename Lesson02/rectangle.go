package main

import (
	"fmt"
)

func main() {
	var width, height int
	fmt.Print("Please enter width and height of the rectangle ")
	fmt.Scanf("%d%d\n", &width, &height)
	fmt.Printf("Area of the rectangle is equal %d\n", 2*(width*height))
}
