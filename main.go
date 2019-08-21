package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
)

var source []byte
var sourceIndex int = 0

func getChar() (byte, error) {
	if sourceIndex == len(source) {
		return 0, errors.New("EOF")
	}
	char := source[sourceIndex]
	sourceIndex++
	return char, nil
}

func ungetChar() {
	sourceIndex--
}

type Token struct {
	kind  string // "intliteral"
	value string
}

func main() {
	var source []byte
	source, _ = ioutil.ReadFile("/dev/stdin")
	var input string = string(source)

	number, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("  .global main\n")
	fmt.Printf("main:\n")
	fmt.Printf("  movq $%d, %%rax\n", number)
	fmt.Printf("  ret\n")
}
