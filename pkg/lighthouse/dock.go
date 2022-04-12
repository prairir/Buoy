/*
	lighthouses have docks, I guess

	has helper functions that hash passwords, get elements from a group from hashes, etc.
*/
package lighthouse

import (
	"crypto/sha512"
)

func NameToHash(networkName string) []byte {
	h := sha512.New()
	h.Write([]byte(networkName))
	return h.Sum(nil)
}

// gets the "distance" between two hashes. This is simply the xor of all the bytes, into a byte array.
func Dist(hash1 []byte, hash2 []byte) []byte {
	for i := range hash1 {
		hash1[i] ^= hash2[i]
	}
	return hash1
}
