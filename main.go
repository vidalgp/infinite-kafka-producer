package main

import (
	"fmt"
	"os"
)

const KEY string = "0123456789ABCDEF0123456789ABCDEF"
const IV string = "0123456789ABCDEF"

func main() {
	fmt.Println("filepath")
	filepath := os.Args[1]
	fmt.Println(filepath)

	lines, _ := readFile(filepath)

	for _, l := range lines {
		fmt.Println(l)
	}

	encrypted := encryptAES([]byte(KEY), []byte(IV), lines)
	encoded := encodeBase64(encrypted)

	produceToKafka(encoded)
}
