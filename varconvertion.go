package main

import "fmt"

func main() {


var i int = 10
var f float64 = 10.1234

fmt.Printf(" Type of I is : %T value is %d \n", i,i)
fmt.Printf(" Type of I is : %T value is %f \n", f,f)

temp := i
i = int (f)
f = float64(temp)

fmt.Printf(" \n Type of I is : %T value is %d \n", i,i)
fmt.Printf(" Type of I is : %T value is %f ", f,f)

}
