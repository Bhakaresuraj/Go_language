package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops in Go language")

	// classic for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// For Range
	// var service []String
	var service []string
	service = []string{"Service-1", "Service-2", "Service-3", "Service-4", "Service-5"}

	for _, svc := range service {
		fmt.Print(svc ," ->")
	}

	//  For as a while
	fmt.Println("\nFor as a while loop .......")
	i:=0

	for i< 10{
		fmt.Print(i,"\n");
		i++;
	}
}
