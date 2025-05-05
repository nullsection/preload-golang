package main

import (
	"C"
	"fmt"
	"log"
	"net"
	"os/exec"

	"github.com/rainycape/dl"
)

var printOnce bool

// main does nothing, just needed for shared library
func main() {}

//export strrchr
func strrchr(s *C.char, c C.int) *C.char{
	// Only print once
	if !printOnce {
		fmt.Println("Hello, world!")
		printOnce = true
	}

	// Open the libc library dynamically
	lib, err := dl.Open("libc", 0)
	if err != nil {
		log.Fatalln(err)
	}
	defer lib.Close()

	// Find the original strrchr function from libc
	var old_strrchr func(s *C.char, c C.int) *C.char
	if err := lib.Sym("strrchr", &old_strrchr); err != nil {
		log.Fatalln("Error retrieving strrchr:", err)
	}

	// Correctly call the old_strrchr function and return its result
	return old_strrchr(s, c)
}
