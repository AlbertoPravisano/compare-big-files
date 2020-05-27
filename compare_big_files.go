package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/sergi/go-diff/diffmatchpatch"
)

const chunkSize = 64000

func main() {

	file1 := os.Args[1]
	file2 := os.Args[2]

	fmt.Println("Start!")

	deepCompare(file1, file2)

	fmt.Println("Done!")
}

func deepCompare(file1, file2 string) bool {
	dmp := diffmatchpatch.New()

	// Check file size ...

	f1, err := os.Open(file1)
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()

	f2, err := os.Open(file2)
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

	for {
		b1 := make([]byte, chunkSize)
		_, err1 := f1.Read(b1)

		b2 := make([]byte, chunkSize)
		_, err2 := f2.Read(b2)

		if err1 != nil || err2 != nil {
			if err1 == io.EOF && err2 == io.EOF {
				return true
			} else if err1 == io.EOF || err2 == io.EOF {
				return false
			} else {
				log.Fatal(err1, err2)
			}
		}

		diffs := dmp.DiffMain(b1, b2, false)

		if !bytes.Equal(b1, b2) {
			fmt.Println(dmp.DiffPrettyText(diffs))
			return false
		}
	}
}
