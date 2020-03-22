package main

import (
	"fmt"
	"strings"

	"github.com/gnyman/go-test-cache-experiment/lib"
)

func main() {
	fmt.Printf("Hello, playground %s\n", lib.Bar())
}

func foo(x string) string {
	return strings.ToUpper(x)
}
