// Package main implements a CLI tool for listing Docker Hub image tags.
package main

import (
	"fmt"
	"os"
	"strings"

	dockerhubimagetags "cli-tools/internal/docker-hub-image-tags"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <namespace> <repository>\n", os.Args[0])
		os.Exit(1)
	}

	namespace := os.Args[1]
	repository := os.Args[2]

	tags, err := dockerhubimagetags.FetchTags(namespace, repository)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(strings.Join(tags, "\n"))
}
