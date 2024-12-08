package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Not enough arguments!")
	}
	var read, valid, nInts, nFloats int
	invalid := make([]string, 0)

	for _, v := range arguments[1:] {
		read++
		_, err := strconv.Atoi(v)
		if err == nil {
			valid++
			nInts++
			continue
		}
		_, err = strconv.ParseFloat(v, 64)
		if err == nil {
			valid++
			nFloats++
			continue
		}
		invalid = append(invalid, v)
	}

	fmt.Printf("#read: %d #valid: %d #int: %d #float: %d\n", read, valid, nInts, nFloats)
	if len(invalid) >= valid {
		fmt.Println("Too many invalid arguments: ", len(invalid))
		for _, iv := range invalid {
			fmt.Println(iv)
		}
	}
}
