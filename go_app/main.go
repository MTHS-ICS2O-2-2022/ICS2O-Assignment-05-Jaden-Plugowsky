// Created by: Jaden Plugowsky
// Created on: May 2023
//
// This program takes a user-given number and calculates the answer based on a pattern

package main

import "fmt"

var layer int
var addby = 3
var answer int // starts at 0
var counter int

func main() {
	// This function takes a user-given number and calculates the answer based on a pattern
	// Input
	fmt.Println("\nThis program takes a user-given number and calculates the answer based on a pattern.")
	fmt.Println("Layer 1: 3, Layer 2: 6, Layer 3: 10, Layer 4: 15, and so on.")
	fmt.Println("\nEnter the layer number you would like to calculate to: ")
	fmt.Scanln(&layer)

	// Process
	for counter < layer {
		answer += addby
		if answer >= 6 {
			addby++
		}
		counter++
	}

	// Output
	fmt.Println("\nLayer", layer, "has", answer, "dots.")

	fmt.Println("\nDone.")
}
