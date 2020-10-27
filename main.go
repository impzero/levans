package main

import "fmt"

func main() {
	n := min(3, 5)
	m := max(4, 5)

	fmt.Println(n)
	fmt.Println(m)

}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}