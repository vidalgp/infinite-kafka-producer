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

type aesPass struct {
	key []byte
	iv  []byte
}

type aesOperator struct {
	pass  *aesPass
	block cipher.Block
}

func NewAesOperator(key []byte, iv []byte) (encrypter, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	operator := aesOperator{
		pass:  &aesPass{key: key, iv: iv},
		block: block,
	}

	return operator, nil
}

func (e aesOperator) encrypt(text string) string {
	encrypter := cipher.NewCBCEncrypter(e.block, e.pass.iv)
	bytetext := []byte(text)
	paddedtext := PKCS5Padding(bytetext, aes.BlockSize)
	ciphertext := make([]byte, len(paddedtext))
	encrypter.CryptBlocks(ciphertext, paddedtext)
	encrypted := base64.StdEncoding.EncodeToString(ciphertext)

	return encrypted
}

func (e aesOperator) decrypt(text string) string {
	decrypter := cipher.NewCBCDecrypter(e.block, e.pass.iv)
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

	decrypter.CryptBlocks(ciphertext, ciphertext)
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
