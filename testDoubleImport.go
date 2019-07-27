/*
	DoubleImport tests somewhat odd syntax for package loading in golang
*/

package main

import "fmt"
import "math"

func testDoubleImport() {

	fmt.Println("Hello World!")
	fmt.Println(math.Abs(-1))
}
