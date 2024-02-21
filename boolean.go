package main

import (
	"fmt"
)

var i, j, k bool

func main() {

	var v int = 2
	var a, b int

	a, b = sq(v)

	fmt.Println("Square Root and cube root is : ", a, b)

}

func sq(xyz int) (x, y int) {

	x = xyz * xyz
	y = xyz * xyz * xyz

	return
}
