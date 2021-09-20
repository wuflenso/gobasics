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
 * good source: https://www.youtube.com/watch?v=sTFJtxJXkaY and https://medium.com/@meeusdylan/when-to-use-pointers-in-go-44c15fe04eac
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

/*
 * BELOW IS A PART TO PROVE WHY TO USE POINTERS IN STRUCT-RELATED FUNCTIONS
 * OVER NO-POINTER
 */
type person struct {
	name string
}

// example of modifying struct value without pointer
func WithoutPointer() {
	p := person{"Richard"}
	p = renameNoPointer(p)
	fmt.Println(p)
}
func renameNoPointer(p person) person {
	p.name = "test"
	return p
}

// example of modifying struct value with pointer (less clutter and redundancy)
func WithPointer() {
	p := person{"Richard"}
	renamePointer(&p)
	fmt.Println(p)
}
func renamePointer(p *person) {
	p.name = "test"
}
