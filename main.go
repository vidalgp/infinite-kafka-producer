package main

import (
	"fmt"
	"os"
)

const KEY string = "1234567890ABCDEF1234567890ABCDEF"
const IV string = "1234567890ABCDEF"

func main() {
	aesOperator, _ := NewAesOperator(KEY, IV)

	fmt.Println("filepath")
	filepath := os.Args[1]
	fmt.Println(filepath)

	lines, _ := readFile(filepath)

	for _, l := range lines {
		fmt.Println(l, "~")
	}

	encrypted := aesOperator.encrypt(lines[0])

	produceToKafka(encrypted)
}
