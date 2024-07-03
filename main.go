package main

import (
	"fmt"

	// Import the random package.
	"github.com/utsabc/neet-code-150/random"
)

func main() {
	// Call the random.Number() function to get the random number. Notice that
	// we use the package name as the accessor, just like we do for the standard
	// library packages.
	fmt.Printf("Your lucky number is %d!\n", random.Number())
}
