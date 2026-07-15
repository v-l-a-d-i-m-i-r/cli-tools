// Package main implements the tuxi CLI entrypoint.
package main

import (
	"cli-tools/internal/tuxi"
	"os"
)

func main() {
	if err := tuxi.NewRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
