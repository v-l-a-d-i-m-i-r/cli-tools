// Package main implements a CLI tool for generating work item IDs.
package main

import (
	"fmt"
	"time"
)

func main() {
	id := time.Now().UTC().Format("20060102150405")
	fmt.Println(id)
}
