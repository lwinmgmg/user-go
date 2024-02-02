package hashing

import (
	"crypto/sha256"
	"encoding/base64"
)

func Hash256(input string) ([]byte, error) {
	sha := sha256.New()
	if _, err := sha.Write([]byte(input)); err != nil {
		return nil, err
	}
	return sha.Sum(nil), nil
}

func Hash256Hex(input string) (string, error) {
	val, err := Hash256(input)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(val), nil
}
