package main

import (
	"fmt"

	"github.com/dafahfabiansyah/go-file-parser/parser"
)

func main() {
	// create file
	fmt.Println("file has created", parser.CreateNewFile("test.txt", "halo dunia"))
	// read file
	fmt.Println(parser.ReadFile("test.txt"))
	// update file
	fmt.Println(parser.UpdateFile("test.txt", "halo dunia, aku atmin"))
	// delete file
	fmt.Println(parser.DeleteFile("test.txt"))
}
