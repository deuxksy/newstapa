package main

import (
	"fmt"
	"math"
)

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string, string) {
	return y, x, x+y
}

func split (sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("Hello", math.Pi, "World")
	fmt.Println("Now you have %g problems.", math.Nextafter(2,3))
	fmt.Println(add(4,2))
	a, b, c :=  swap("hello", "world")
	fmt.Println(a, b, c)
	fmt.Println(split(17))
}
