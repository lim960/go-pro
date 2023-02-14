package util

import (
	"crypto/rand"
	uuid "github.com/satori/go.uuid"
	"math/big"
	"strings"
)

func Get32() string {
	return strings.ReplaceAll(uuid.NewV4().String(), "-", "")
}

func GetRandStr(n int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")
	result := make([]byte, n)
	for i := range result {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		result[i] = letters[n.Int64()]
	}

	return string(result)
}

func GetFileName(n int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyz123456789")
	result := make([]byte, n)
	for i := range result {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		result[i] = letters[n.Int64()]
	}

	return string(result)
}
