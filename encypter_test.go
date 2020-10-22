package main

import (
	"testing"
)

type testTable struct {
	value    string
	expected string
}

func TestNewAesOperator(t *testing.T) {
	iv := "0123456789ABCDEF"
	key := "0123456789ABCDEF0123456789ABCDEF"

	_, err := NewAesOperator(key, iv)
	if err != nil {
		t.Errorf("Expected valid instantiation, instead got error %v", err)
	}

	_, err = NewAesOperator("invalidkey", iv)
	if err == nil {
		t.Errorf("Expected error on invalid key")
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic on invalid iv")
		}
	}()
	NewAesOperator(key, "invalid iv")
}

func TestEncrypt(t *testing.T) {
	var testkey string = "0123456789ABCDEF0123456789ABCDEF"
	var testiv string = "0123456789ABCDEF"
	tests := []testTable{
		testTable{"hello world", "XdkOMEi7cNQ4j9Baiw3UmQ=="},
		testTable{"test123456789012", "e5cwjRQoq6C2YIFhpLi8zKm7ATuiBjnBYJZ2OtaVoDw="},
		testTable{"", "/q4TRsmg40xXGfIjOG6pqg=="},
	}

	aesOperator, _ := NewAesOperator(testkey, testiv)
	for _, v := range tests {
		a := aesOperator.encrypt(v.value)
		if a != v.expected {
			t.Errorf("Expected '%v', but got %v", v.expected, a)
		}
	}

}

func TestDecrypt(t *testing.T) {
	var testkey string = "0123456789ABCDEF0123456789ABCDEF"
	var testiv string = "0123456789ABCDEF"
	tests := []testTable{
		testTable{"XdkOMEi7cNQ4j9Baiw3UmQ==", "hello world"},
		testTable{"e5cwjRQoq6C2YIFhpLi8zKm7ATuiBjnBYJZ2OtaVoDw=", "test123456789012"},
		testTable{"/q4TRsmg40xXGfIjOG6pqg==", ""},
	}

	aesOperator, _ := NewAesOperator(testkey, testiv)
	for _, v := range tests {
		a := aesOperator.decrypt(v.value)
		if a != v.expected {
			t.Errorf("Expected '%v', but got %v", v.expected, a)
		}
	}

}
