package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: go run . <input> <output>")
		os.Exit(1)
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	result := Process(string(data))
	if err := os.WriteFile(os.Args[2], []byte(result), 0o644); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
