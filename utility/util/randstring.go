package util

import (
	"math/rand"
	"time"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const numset = "0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func GetRandomString(length int) string {

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GetRandomNumber(length int) string {

	b := make([]byte, length)
	for i := range b {
		b[i] = numset[seededRand.Intn(len(numset))]
	}
	return string(b)
}