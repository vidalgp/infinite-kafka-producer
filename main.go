package main

import (
	"os"
)

const KEY string = "1234567890ABCDEF1234567890ABCDEF"
const IV string = "1234567890ABCDEF"
const TOPIC string = "test-topic-a"

func main() {
	filepath := os.Args[1]

	lines, _ := readFile(filepath)

	aesOperator, _ := NewAesOperator([]byte(KEY), []byte(IV))
	var encryptedMessages []string
	for _, v := range lines {
		encrypted := aesOperator.encrypt(v)
		encryptedMessages = append(encryptedMessages, encrypted)
	}

	produceToKafka(TOPIC, encryptedMessages...)
}
