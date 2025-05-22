package crypto

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"io"
)

var randReader io.Reader = rand.Reader

// DoubleSHA256 returns SHA256(SHA256(data)).
func DoubleSHA256(data []byte) []byte {
	first := sha256.Sum256(data)
	second := sha256.Sum256(first[:])
	return second[:]
}

// HMACSHA256 computes HMAC-SHA256 over data using key.
func HMACSHA256(key, data []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	return h.Sum(nil)
}

// RandomBytes returns cryptographically secure random bytes of length n.
func RandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := io.ReadFull(randReader, b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
