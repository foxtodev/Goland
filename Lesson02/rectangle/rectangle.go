package rectangle

import (
	"fmt"
)

func Area() {
	var width, height int
	fmt.Print("Please enter size of rectangle (width, height)")
	fmt.Scanf("%d%d\n", &width, &height)
	fmt.Printf("Area of the rectangle is equal %d\n", width*height)
}
