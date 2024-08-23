package main

import (
	"fmt"

	"github.com/guionardo/gs-ops/internal/commons"
)

func main() {
	fmt.Printf("%s CLI v%s", commons.AppName, commons.Version)
}
