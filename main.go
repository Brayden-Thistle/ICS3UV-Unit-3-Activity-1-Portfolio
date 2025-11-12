/*
* @author Brayden Thistle
* @version 1.0.0
* @date 2025-11-11
* @fileoverview this program displays and adds 3 numbers.
*/

package main

import "fmt"

func main() {
	// the numbers as constants
	const count int = 10 
	const count1 int = -20
	const count2 int = 85

	// calculating the sum
	answer := count + count1 + count2  

//displaying the answer
	fmt.Println("The sum of ", count, "and", count1, "and", count2, "equals", answer)

	fmt.Println ("\nDone")
}