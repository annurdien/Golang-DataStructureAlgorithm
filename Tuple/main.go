package main

import "fmt"

// Tuple di Go

// fungsi return 2 integer yaitu square dari a dan cube dari a

func powerSeries(a int) (int, int) {

	// Return suqare, dan cube
	return a * a, a * a * a

}

// Same but different implementation
/*
func powerSeries(a int) (square int, cube int) {

	square = a * a
	cube = a * a * a
	return
}

*/

/*
func powerSeries(a int) (int, int, error) {

	var square = a * a
	var cube = square * a

	return square, cube, nil
}
*/
func main() {

	var square int
	var cube int

	// store return value dari power series
	square, cube = powerSeries(3)

	fmt.Println("Square ", square, "Cube ", cube)

}
