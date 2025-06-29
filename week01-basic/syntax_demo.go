package main

import "fmt"

func main() {
	var x = 10
	y := 20
	fmt.Println("x + y = ", add(x, y))

	names := []string{"Go", "Rust", "Python"}
	for i, name := range names {
		fmt.Printf("%d: %s\n", i, name)
	}

	m := map[string]int{"apple": 1, "banana": 2}
	fmt.Println("banana : ", m["banana"])
	fmt.Println("cherry : ", m["cherry"])
}

func add(x int, y int) int {
	return x + y
}
