package utils

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	mrand "math/rand"
	"time"
)

func GenerateRandomID() string {
	b := make([]byte, 4) // 4 bytes == 8 hex characters
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(b)
}

func ShuffleSlice(slice []string) {
	mrand.New(mrand.NewSource(time.Now().UnixNano()))
	mrand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}
