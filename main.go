package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 1) looping for all file
	for _, fileName := range os.Args[1:] {
		// 2) open the file
		fileHandle, err := os.Open(fileName)

		// if err not nil print the err and continue to another file
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		// 3) copy the file to Stdout

		if _, err := io.Copy(os.Stdout, fileHandle); err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		// 4) close the file
		fileHandle.Close()
	}
}
