package main

import (
	"fmt"
)

func main() {

	mySlice := make([]int, 1, 4)
	fmt.Printf("Lenght starts out as %d with a capacity of %d\n",
		len(mySlice), cap(mySlice))
		
	for i := 1; i < 17; i++ {
		mySlice = append(mySlice, i)
		fmt.Printf("Slice lenght is %d but capacity is %d\n",
			len(mySlice), cap(mySlice))
	}
}