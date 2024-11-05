package utils

import (
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/sha3"
)

func SHA3Hash(input string) string {
	hash := sha3.New256()
	_, _ = hash.Write([]byte(input))
	sha3 := hash.Sum(nil)
	return fmt.Sprintf("%x", sha3)
}

func GetUUID() string {
	return uuid.New().String()
}
