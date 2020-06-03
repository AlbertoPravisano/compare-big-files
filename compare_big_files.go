package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jwangsadinata/go-multimap/slicemultimap"
)

const fromBtoMB = 1024 * 1024

func main() {
	deepCompare(os.Args[1], os.Args[2])
	fmt.Println("Done!")
}

func deepCompare(file1, file2 string) bool {

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

	newm := slicemultimap.New()

	row1 := 1
	row2 := 1

	//LOAD CHUNCK
	scanner := bufio.NewScanner(f1)
	buf := make([]byte, 0, 1*fromBtoMB)
	scanner.Buffer(buf, 10*fromBtoMB)

	//LOAD MAP WITH FILE1 ROWS
	for scanner.Scan() {
		newm.Put(scanner.Text(), "File 1 - Row "+strconv.Itoa(row1))
		row1++
	}

	//COMPARE CHUNCKS
	scanner2 := bufio.NewScanner(f2)
	buf2 := make([]byte, 0, 1*fromBtoMB)
	scanner2.Buffer(buf2, 10*fromBtoMB)

	//KEY:ROW_STRING
	//VALUE:"File x - Row y"

	//DELETE MATCHING ROWS FROM FILE2
	for scanner2.Scan() {
		val, found := newm.Get(scanner2.Text())
		if found {
			newm.Remove(scanner2.Text(), val[0])
		} else {
			newm.Put(scanner2.Text(), "File 2 - Row "+strconv.Itoa(row2))
		}
		row2++
	}

	//PRINT NOT MATCHED ROWS
	fmt.Println("There are", newm.Size(), "differencies")
	for _, key := range newm.KeySet() {
		value, _ := newm.Get(key)
		fmt.Println(value, "-", key)
	}

	return true
}
