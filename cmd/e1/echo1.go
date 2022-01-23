package main

import (
	"fmt"
	"os"
)

func main() {
	for i, s := range os.Args[:] {
		fmt.Printf("Index: %v\tValue: %v\n", i, s)
	}
}
