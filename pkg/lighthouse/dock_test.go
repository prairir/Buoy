package lighthouse

import (
	"testing"
)

func TestNameToHash(t *testing.T) {
	hash1 := NameToHash("test")
	hash2 := NameToHash("tess")
	t.Log(Dist(hash1, hash2))
}
