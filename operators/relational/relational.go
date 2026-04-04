package main

import "fmt"

func main() {
	a, b := 20, 10

	fmt.Println("--- Relational Operations ---")
	fmt.Printf("Is %d == %d? %v\n", a, b, a == b)
	fmt.Printf("Is %d != %d? %v\n", a, b, a != b)
	fmt.Printf("Is %d > %d?  %v\n", a, b, a > b)

	fmt.Println("\n--- Logical Operations ---")
	isAdult := true
	hasLicense := false

	// Logical AND (දෙකම true විය යුතුයි)
	fmt.Println("Can Drive (AND):", isAdult && hasLicense)

	// Logical OR (එකක් true වීම සෑහේ)
	fmt.Println("Either is True (OR):", isAdult || hasLicense)

	// Logical NOT (අගය මාරු කරයි)
	fmt.Println("Reverse isAdult (NOT):", !isAdult)
}
