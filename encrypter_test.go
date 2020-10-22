package main

import (
	"testing"
)

type testTable struct {
	value    string
	expected string
}

func BenchmarkEncrypt(b *testing.B) {
	key := []byte("0123456789ABCDEF0123456789ABCDEF")
	iv := []byte("0123456789ABCDEF")
	aesOperator, _ := NewAesOperator(key, iv)
	input := "hello world"

	for n := 0; n < b.N; n++ {
		encrypted := aesOperator.encrypt(input)
		aesOperator.decrypt(encrypted)
	}
}

func TestNewAesOperator(t *testing.T) {
	iv := []byte("0123456789ABCDEF")
	key := []byte("0123456789ABCDEF0123456789ABCDEF")

	_, err := NewAesOperator(key, iv)
	if err != nil {
		t.Errorf("Expected valid instantiation, instead got error %v", err)
	}

	_, err = NewAesOperator([]byte("invalidkey"), iv)
	if err == nil {
		t.Errorf("Expected error on invalid key")
	}

}

func TestEncrypt(t *testing.T) {
	iv := []byte("0123456789ABCDEF")
	key := []byte("0123456789ABCDEF0123456789ABCDEF")
	tests := []testTable{
		testTable{"hello world", "XdkOMEi7cNQ4j9Baiw3UmQ=="},
		testTable{"test123456789012", "iKZXie3y/3zC+OJiSQnjgYs4dHpX73sadEEJpDpUxhY="},
		testTable{"", "hOrDZZmk9Osg6/YFJM47Mw=="},
	}

	aesOperator, _ := NewAesOperator(key, iv)
	for _, v := range tests {
		a := aesOperator.encrypt(v.value)
		if a != v.expected {
			t.Errorf("Expected '%v', but got %v", v.expected, a)
		}
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic on invalid iv")
		}
	}()
	aesOperator, _ = NewAesOperator(key, []byte("invalid iv"))
	aesOperator.encrypt("hello world")
}

func TestDecrypt(t *testing.T) {
	iv := []byte("0123456789ABCDEF")
	key := []byte("0123456789ABCDEF0123456789ABCDEF")
	tests := []testTable{
		testTable{"XdkOMEi7cNQ4j9Baiw3UmQ==", "hello world"},
		testTable{"iKZXie3y/3zC+OJiSQnjgYs4dHpX73sadEEJpDpUxhY=", "test123456789012"},
		testTable{"hOrDZZmk9Osg6/YFJM47Mw==", ""},
	}

	aesOperator, _ := NewAesOperator(key, iv)
	for _, v := range tests {
		a := aesOperator.decrypt(v.value)
		if a != v.expected {
			t.Errorf("Expected '%v', but got %v", v.expected, a)
		}
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic on invalid iv")
		}
	}()
	aesOperator, _ = NewAesOperator(key, []byte("invalid iv"))
	aesOperator.decrypt("XdkOMEi7cNQ4j9Baiw3UmQ==")

}
