package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Get Home variable path
var home = os.Getenv("HOME")
var err error

// Get windows path of `$HOME/../../`
var slash = filepath.Join(home, "./..", "./..")

func main() {

	// Tell it how to tell error
	slash, err := filepath.Abs(slash)
	if err != nil {
		log.Fatal(err)
	}
	
	// Print Install Dir
	fmt.Println("Install Dir:", slash)
}
