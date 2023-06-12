package main

import (
	"fmt"
)

func main() {
	bestFinish := championshipFinishes(12, 5, 6, 4, 3, 3)

	fmt.Println(bestFinish)
}

func championshipFinishes(finishes ...int) int{
	best := finishes[0]
	for _, i := range finishes {
		if i < best {
			best = i
		}
	}
	return best
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	// Accepting return value in varaible
// 	sum := add(20, 30)
// 	fmt.Println(sum)
// }
// // Function with int as return type
// func add(x int, y int) int {
// 	total := 0
// 	total = x + y
// 	return total
// }
