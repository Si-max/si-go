package main

import (
	"fmt"
	"reflect"
	"strconv"
)


func main() {
	name := "Simon Usifoh" 
	course := "Getting Started with Kubernetes"
	module := "4" //Needs to be integer
	clip := 2 //Needs to be interger
	// courseComplete := false

	fmt.Println("Name and Course are set to", name, "and", course, ".")
	fmt.Println("Module and Clip are set to", module, "and", clip, ".")
	fmt.Println("Name is of type", reflect.TypeOf(name))
	fmt.Println("Module is of type", reflect.TypeOf(module))
	// total := module + clip
	// fmt.Println("Module plus clip equals", total)
	
	iModule, err := strconv.Atoi(module)
	if err == nil {
		total := iModule + clip
		fmt.Println("Module plus clip equals", total)
	}
	fmt.Println("Memory address of *course* variable is", &course )

	var ptr  *string = &course
	fmt.Println("Pointing course variable at address,", ptr, "which holds this value,", *ptr)
}
