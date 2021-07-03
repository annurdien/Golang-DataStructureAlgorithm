package main

import "fmt"

// fungsi return 2 integer yaitu square dari a dan cube dari a

func powerSeries(a int) (int, int) {

	// Return suqare, dan cube
	return a * a, a * a * a

}

func main() {

	var square int
	var cube int

	// store return value dari power series
	square, cube = powerSeries(3)

	fmt.Println("Square ", square, "Cube ", cube)

}
