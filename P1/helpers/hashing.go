package helpers

import "crypto/sha256"

func HashSHA256(a string) []byte {
	h := sha256.New()
	b := []byte(a)
	hashed := h.Sum(b)
	return hashed
}
