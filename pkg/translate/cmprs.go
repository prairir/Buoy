package translate

// Ryan thought the name was v funny

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
)

// translate.compress: compresses plainText with gzip
func compress(plainText []byte) ([]byte, error) {
	buf := bytes.Buffer{}

	// use the buffer as the memory to write to
	writer := gzip.NewWriter(&buf)

	_, err := writer.Write(plainText)
	if err != nil {
		return []byte{}, fmt.Errorf("translate.compress: %w", err)
	}

	err = writer.Close() // flush the cache
	if err != nil {
		return []byte{}, fmt.Errorf("translate.compress: %w", err)
	}

	return buf.Bytes(), nil
}

// translate.compress: decompresses compressedText with gzip
func decompress(compressedText []byte) ([]byte, error) {
	buf := bytes.Buffer{}

	compressedTextReader := bytes.NewReader(compressedText)
	reader, err := gzip.NewReader(compressedTextReader)
	if err != nil {
		return []byte{}, fmt.Errorf("translate.decompress: %w", err)
	}

	_, err = io.Copy(&buf, reader)
	if err != nil {
		return []byte{}, fmt.Errorf("translate.decompress: %w", err)
	}

	return buf.Bytes(), nil
}
