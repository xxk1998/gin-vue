package util

import (
	"math/rand"
	"time"
)

func RandomString(n int) string {
	var letters = []byte("WEFIHBWEJHEWFBIUfbnskfjbwefbe")
	res := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range res {
		res[i] = letters[rand.Intn(len(letters))]
	}
	return string(res)
}
