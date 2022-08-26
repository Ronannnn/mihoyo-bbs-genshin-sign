package util

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// GetRandString Get a random string with a fixed length
// Reference: https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func GetRandString(length int) string {
	bytes := make([]rune, length)
	for i := range bytes {
		bytes[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(bytes)
}
