package main

import (
	"fmt"

	"github.com/dafahfabiansyah/go-file-parser/parser"
)

func main() {
	// create file
	fmt.Println("file has created", parser.CreateNewFile("file.txt", "hello world"))
	// read file
	fmt.Println(parser.ReadFile("file.txt"))
	// update file

	// delete file
}
