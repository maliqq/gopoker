package util

import (
	"code.google.com/p/go.crypto/bcrypt"
)

import (
	"log"
)

const delimiter = ""
const cost = bcrypt.DefaultCost

func HashPasswordWithSalt(password, salt string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(passwordWithSalt(password, salt)), cost)
	if err != nil {
		log.Fatalf("hashing function error: %s", err)
	}
	return string(hash)
}

func CompareHashAndPasswordWithSalt(hash, password, salt string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwordWithSalt(password, salt)))
}

func passwordWithSalt(password string, salt string) string {
	return password + delimiter + salt
}
