package main

import "fmt"

func main() {
	sum := 0

	for i := 1; i <= 100; i++ {
		sum += i
		fmt.Printf("%3d: %4d\n", i, sum)
	}
}
