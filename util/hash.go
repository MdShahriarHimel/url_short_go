package util

import (
	"crypto/sha256"
)

func Hash(data string) *string {
	h := sha256.New()
	_, err := h.Write([]byte(data))
	if err != nil {
		return nil
	}

	byteHash := h.Sum(nil)

	strHash := string(byteHash)
	return &strHash
}
