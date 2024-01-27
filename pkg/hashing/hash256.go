package hashing

import "crypto/sha256"

func Hash256(input string) ([]byte, error) {
	sha := sha256.New()
	if _, err := sha.Write([]byte(input)); err != nil {
		return nil, err
	}
	return sha.Sum(nil), nil
}
