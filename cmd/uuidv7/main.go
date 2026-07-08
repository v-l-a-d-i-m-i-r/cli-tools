// Package main implements a CLI tool for generating UUIDv7 values.
package main

import (
	"fmt"
	"os"

	"github.com/google/uuid"
)

func main() {
	id, err := uuid.NewV7()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(id)
}
