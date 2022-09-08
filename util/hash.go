package util

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
)

func HashFile(file string) string {
	hasher := sha256.New()
	reader, err := os.ReadFile(file)
	hasher.Write(reader)
	if err != nil {
		return ""
	}
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash
}

func HashString(str string) string {
	hasher := sha256.New()
	hasher.Write([]byte(str))
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash
}
