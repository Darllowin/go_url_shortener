package shortener

import (
	"math/rand/v2"
	"strings"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateKey(n int) string {
	var genKey strings.Builder

	for range n {
		genKey.WriteString(string(charset[rand.N(len(charset))]))
	}

	return genKey.String()
}
