package main

import (
	"fmt"
	"io"
	"os"
)

// go run copyfile.go source target
func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: copyfile <src> <dst>")
		return
	}

	sourceFile := os.Args[1]
	destFile := os.Args[2]

	src, err := os.Open(sourceFile)
	if err != nil {
		fmt.Printf("open source file fatal: %v\n", err)
	}
	defer func(src *os.File) {
		err := src.Close()
		if err != nil {
			fmt.Printf("close source file fatal: %v\n", err)
		}
	}(src)

	// create a target file
	det, err := os.Create(destFile)
	if err != nil {
		fmt.Printf("create dest file fatal: %v\n", err)
		return
	}
	defer func(det *os.File) {
		err := det.Close()
		if err != nil {
			fmt.Printf("close dest file fatal: %v\n", err)
		}
	}(det)

	n, err := io.Copy(det, src)
	if err != nil {
		fmt.Printf("copy file fatal: %v\n", err)
		return
	}

	fmt.Printf("copy file completed! copied %d bytes\n", n)
}
