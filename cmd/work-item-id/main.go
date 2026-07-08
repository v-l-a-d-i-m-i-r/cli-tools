package main

import (
	"fmt"
	"time"
)

func main() {
	id := time.Now().UTC().Format("20060102150405")
	fmt.Println(id)
}
