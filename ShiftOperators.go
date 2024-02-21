package main

import "fmt"

func main() {

	var left uint = 1 << 9
	fmt.Printf("\n Left Shift value is : %d", left)
	left = left >> 8
	fmt.Printf("\n Right Shift value is : %d", left)
	left = 1
	fmt.Printf("\n %d", left<<1)
	fmt.Printf("\n %d", left<<1)
	fmt.Printf("\n %d", left<<2)
	fmt.Printf("\n %d", left<<3)
	fmt.Printf("\n %d", left<<4)
	fmt.Printf("\n %d", left<<5)
	fmt.Printf("\n %d", left<<6)
	fmt.Printf("\n %d", left<<7)
	fmt.Printf("\n %d", left<<8)
	fmt.Printf("\n %d", left<<9)
	fmt.Printf("\n %d", left<<10)
	fmt.Printf("\n %d", left<<11)

}
