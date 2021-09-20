package basic

import (
	"fmt"
)

/* POINTERS
 * Keys to understand pointers:
 * 1. Ampersand (&) creates a pointer object	=> &s
 * 2. Asterisk (*) shows the pointer value		=> *s
 *
 * The thing that usually fucks up newbs is the fact that
 * pointer type is written as *int (see Example 1) while *variablename
 * shows the pointer's stored value. In this case variablename will
 * show the memory address.
 * good source: https://www.youtube.com/watch?v=sTFJtxJXkaY
 *
 * Using := in your code is much better to understand pointers
 * (see Example 2)
 */

func RunPointers() {

	// Example 1
	var num int = 5
	var powint *int = &num // *int is different from *powint
	fmt.Println(*powint)

	// Example 2
	var text string = "where u at"
	pointa := &text
	fmt.Println(pointa)  // prints the memory address
	fmt.Println(*pointa) // prints the pointer value
}
