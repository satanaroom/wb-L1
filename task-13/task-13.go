package main

import "fmt"

func main() {
	n := 792341293
	m := -2133962345

	fmt.Printf("numbers before swap: n = %d, m = %d\n", n, m)
	n, m = m, n
	fmt.Printf("numbers after  swap: n = %d, m = %d\n", n, m)
}
