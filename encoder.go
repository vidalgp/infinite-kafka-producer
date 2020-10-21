package main

import (
	"encoding/base64"
	"fmt"
)

func encodeBase64(messages []string) []string {
	var encodedMessages []string
	for _, msg := range messages {
		encoded := base64.StdEncoding.EncodeToString([]byte(msg))
		fmt.Println(encoded)
		encodedMessages = append(encodedMessages, encoded)
	}
	return encodedMessages
}
