// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

//!+
func main() {

	for i, a := range os.Args[:] {
		fmt.Printf("%v, %v\n", i, a)
	}
}

//!-
