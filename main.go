package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s expression\n", os.Args[0])
		os.Exit(1)
	}

	result, err := Evaluate(os.Args[1])
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(0)
	}

	fmt.Printf("%d\n", result)
}





