package main

import (
	"fmt"

	"github.com/KTBL/txt-tool/files"
)

// go run ./main.go
func main() {

	fmt.Printf("voll das crazy zeug hier ... hab ab!! U+2191U+2191\n")

	file := "data/txt-1.txt"

	files.ReadLineByLine(&file)

	fmt.Printf("... ths is the end, my only friend the end!!")

}
