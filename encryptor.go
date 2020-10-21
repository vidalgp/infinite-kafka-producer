package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func encryptAES(key []byte, iv []byte, messages []string) []string {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	mode := cipher.NewCBCEncrypter(block, iv)

	var encryptedMessages []string
	for _, msg := range messages {
		msgbyte := []byte(msg)
		ciphertext := make([]byte, aes.BlockSize+len(msgbyte))
		fmt.Println(len(msg))
		if len(msg)%aes.BlockSize != 0 {
			panic("input text is not a multiple of the block size")
		}
		mode.CryptBlocks(ciphertext[aes.BlockSize:], msgbyte)
		fmt.Printf("%x\n", ciphertext)
		encryptedMessages = append(encryptedMessages, string(ciphertext))
	}
	return encryptedMessages

}
