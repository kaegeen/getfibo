package main

import (
	"fmt"
)

// generateFibonacci generates the Fibonacci sequence up to n terms.
func generateFibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	}

	// Initialize the sequence with the first two Fibonacci numbers
	sequence := []int{0, 1}

	for i := 2; i < n; i++ {
		next := sequence[i-1] + sequence[i-2]
		sequence = append(sequence, next)
	}

	return sequence[:n]
}

func main() {
	var terms int

	fmt.Println("Fibonacci Sequence Generator")
	fmt.Println("============================")
	fmt.Print("Enter the number of terms: ")
	_, err := fmt.Scan(&terms)

	// Handle invalid input
	if err != nil || terms < 1 {
		fmt.Println("Please enter a valid positive integer.")
		return
	}

	// Generate and display the Fibonacci sequence
	sequence := generateFibonacci(terms)
	fmt.Printf("Fibonacci sequence (%d terms): %v\n", terms, sequence)
}
