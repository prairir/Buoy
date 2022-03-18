package translate

import (
	"bytes"
	"testing"
)

func TestCompress(t *testing.T) {
	answer := []byte{31, 139, 8, 0, 0, 0, 0, 0, 0, 255, 202, 72, 205, 201, 201, 87, 40, 201, 72, 45, 74, 5, 4, 0, 0, 255, 255, 80, 233, 18, 109, 11, 0, 0, 0}

	result, err := compress([]byte("hello there"))
	if err != nil {
		t.Log("translate.compress error should be nil:", err)
		t.Fail()
	}

	if !bytes.Equal(result, answer) {
		t.Log("result is not answer")
		t.Fail()
	}
}

func TestDecompress(t *testing.T) {
	answer := []byte("hello there")

	input := []byte{31, 139, 8, 0, 0, 0, 0, 0, 0, 255, 202, 72, 205, 201, 201, 87, 40, 201, 72, 45, 74, 5, 4, 0, 0, 255, 255, 80, 233, 18, 109, 11, 0, 0, 0}

	result, err := decompress(input)
	if err != nil {
		t.Log("translate.decompress error should be nil:", err)
		t.Fail()
	}

	if !bytes.Equal(result, answer) {
		t.Log("result is not answer")
		t.Fail()
	}
}
