package main

import (
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	encodedString := "=?koi8-r?B?9MXT1M/Xz8Ug08/Pwt3FzsnF?= <unencoded@mail.com>"
	decodedString := decodeRFC2047(encodedString)
	fmt.Printf("Encoded: %s, decoded: %s\n", encodedString, decodedString)
	plainString := "Unencoded plain text"
	decodedString = decodeRFC2047(plainString)
	fmt.Printf("Encoded: %s, decoded: %s\n", plainString, decodedString)
}
