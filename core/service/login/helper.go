package login

import (
	rands "crypto/rand"
	"io"
	"math/rand"
	"time"
)

var Numbers = []rune("123456789")
var letters = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

func RandAlphanumeric(n int) string {
	a := make([]byte, n)
	if _, err := io.ReadFull(rands.Reader, a); err != nil {

		return ""
	}
	for i := range a {
		a[i] = letters[int(a[i])%len(letters)]
	}
	return string(a)
}

func GenerateRandomOtp(n int) string {
	if n == 0 {
		n = 6
	}
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = Numbers[rand.Intn(len(Numbers))]
	}
	return string(b)
}
