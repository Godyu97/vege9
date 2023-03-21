package vegeTools

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandStringMask(n int) string {
	b := new(strings.Builder)
	b.Grow(n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return b.String()
}

func RandBytesMask(n int) []byte {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return b
}

// HashBySalt
// hmac sha2 and salt make hash
func HashBySalt(ps, salt string) string {
	mac := hmac.New(sha256.New, []byte(ps))
	mac.Write([]byte(salt))
	hs := mac.Sum(nil)
	return fmt.Sprintf("%x", hs)
}

// CheckBySalt
// check HashBySalt
func CheckBySalt(checked, hash, salt string) bool {
	expectedMAC := HashBySalt(checked, salt)
	return hmac.Equal([]byte(hash), []byte(expectedMAC))
}
