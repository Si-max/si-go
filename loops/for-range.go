package main

import (
	"fmt"
)

func main() {

	courseInProg := []string {
		"Docker & Kubernetes: The Big Picture",
		"Docker Networking",
		"Getting Started with Kubernetes",
		"Kubernetes Deep Dive" }

	courseCompleted := []string {
		"Docker & Kubernetes: The Big Picture",
		"Docker Deep Dive"}
		
	for _, i := range courseInProg {
		fmt.Println(i)
		for _, j := range courseCompleted {
			if i == j {
				fmt.Println("\nHey we found a clash.", i, "is in both lists.")
			}			
		}
	}
}