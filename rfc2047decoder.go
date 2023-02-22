package main

import (
	"bytes"
	"errors"
	"io"
	"mime"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

func decodeRFC2047(rfc2047header string) string {
	dec := mime.WordDecoder{CharsetReader: koi8rSupportingCharsetReader}
	header, err := dec.DecodeHeader(rfc2047header)
	if err != nil {
		panic(err)
	}
	return header
}

// koi8rSupportingCharsetReader is a custom CharsetReader implementation for WordDecoder
func koi8rSupportingCharsetReader(charset string, input io.Reader) (io.Reader, error) {
	switch strings.ToLower(charset) {
	case "koi8-r":
		return koi8rDecoder{input}, nil
	default:
		return nil, errors.New("unsupported charset")
	}
}

type koi8rDecoder struct {
	r io.Reader
}

func (k koi8rDecoder) Read(p []byte) (n int, err error) {
	buf := make([]byte, len(p))
	_, err = k.r.Read(buf)
	if err != nil {
		return 0, err
	}
	copy(p, koi8rToUtf8(buf))
	pos := bytes.IndexByte(p, 0)
	if pos == -1 {
		pos = len(p)
	}
	return pos, nil
}

func koi8rToUtf8(k []byte) []byte {
	// Create a new decoder that will convert KOI8-R to UTF-8
	koi8rDecoder := charmap.KOI8R.NewDecoder()

	// Use the decoder to convert the byte slice from KOI8-R to UTF-8
	u, err := koi8rDecoder.Bytes(k)
	if err != nil {
		panic(err)
	}
	return u
}
