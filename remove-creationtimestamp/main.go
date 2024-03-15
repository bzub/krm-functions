package main

import (
	"fmt"
)

func main() {
	if err := NewCommand().Execute(); err != nil {
		fmt.Printf("error running processor: %s\n", err)
	}
}
