package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

type encrypter interface {
	decrypt(s string) string
	encrypt(s string) string
}

type aesOperator struct {
	encrypter cipher.BlockMode
	decrypter cipher.BlockMode
}

func NewAesOperator(key string, iv string) (encrypter, error) {
	ivec := []byte(iv)
	bkey := []byte(key)

	block, err := aes.NewCipher(bkey)
	if err != nil {
		return nil, err
	}

	operator := aesOperator{
		encrypter: cipher.NewCBCEncrypter(block, ivec),
		decrypter: cipher.NewCBCDecrypter(block, ivec),
	}

	return operator, nil
}

func (e aesOperator) encrypt(text string) string {
	bytetext := []byte(text)

	paddedtext := PKCS5Padding(bytetext, aes.BlockSize)

	ciphertext := make([]byte, len(paddedtext))
	e.encrypter.CryptBlocks(ciphertext, paddedtext)

	encrypted := base64.StdEncoding.EncodeToString(ciphertext)

	return encrypted
}

func (e aesOperator) decrypt(text string) string {
	ciphertext, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		fmt.Println("decode error:", err)
		return ""
	}

	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	e.decrypter.CryptBlocks(ciphertext, ciphertext)

	trimmed := PKCS5Trimming(ciphertext)

	return string(trimmed)
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
