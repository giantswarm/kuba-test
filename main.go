package main

import (
	"fmt"

	"github.com/giantswarm/kuba-test/pkg/project"
)

func main() {
	fmt.Printf("kuba-test: %s\n", project.Version())
}
