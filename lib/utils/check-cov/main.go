package main

import (
	"fmt"
	"os"

	"github.com/ysmood/got"
)

func main() {
	err := got.EnsureCoverage("coverage.txt", 100)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
