package translate

import (
	"bytes"
	"testing"
)

func TestEncryptAndDecrypt(t *testing.T) {
	password := []byte("0123456789012345")
	answer := []byte("hello this is a test")

	resultEnc, err := encrypt(answer, password)
	if err != nil {
		t.Log("translate.encrypt error should be nil:", err)
		t.Fail()
	}

	resultDec, err := decrypt(resultEnc, password)
	if err != nil {
		t.Log("translate.decrypt error should be nil:", err)
		t.Fail()
	}

	if !bytes.Equal(answer, resultDec) {
		t.Log("result is not answer")
		t.Fail()
	}
}
