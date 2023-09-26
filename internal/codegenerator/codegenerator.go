package codegenerator

import (
	"math/rand"
	"time"
)

func codeGenerator(length int) string {
	source := rand.New(rand.NewSource(time.Now().UnixNano()))

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[source.Intn(len(charset))]
	}
	return string(b)
}
