package main

import (
	"fmt"

	"exercism-go/tree"
)

func main() {
	if tree.Ping() {
		fmt.Printf("It's alive\n")
	} else {
		fmt.Printf("It's dead, Jim\n")
	}
}
