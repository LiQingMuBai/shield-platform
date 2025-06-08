package utils

import (
	"log"
	"testing"
)

func TestBcryptHash(t *testing.T) {

	hash := BcryptHash("123456")
	log.Println(hash)
}
