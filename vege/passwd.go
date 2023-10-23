package vege

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var vegeRand *rand.Rand

func init() {
	vegeRand = rand.New(rand.NewSource(time.Now().UnixNano()))
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
	// A vegeRand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, vegeRand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = vegeRand.Int63(), letterIdxMax
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
	// A vegeRand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, vegeRand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = vegeRand.Int63(), letterIdxMax
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

const codeBytes = "0123456789"

func RandCodeMask(n int) string {
	b := new(strings.Builder)
	b.Grow(n)
	// A vegeRand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, vegeRand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = vegeRand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(codeBytes) {
			b.WriteByte(codeBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return b.String()
}

const idBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandIdMask(n int) string {
	b := new(strings.Builder)
	b.Grow(n)
	// A vegeRand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, vegeRand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = vegeRand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(idBytes) {
			b.WriteByte(idBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return b.String()
}

func RandSelfDefMask(n int, bs string) string {
	b := new(strings.Builder)
	b.Grow(n)
	// A vegeRand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, vegeRand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = vegeRand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(bs) {
			b.WriteByte(bs[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return b.String()
}

// HmacHashWithSalt
// hmac sha2 and salt make hash
func HmacHashWithSalt(ps, salt string) string {
	mac := hmac.New(sha256.New, []byte(ps))
	mac.Write([]byte(salt))
	hs := mac.Sum(nil)
	return fmt.Sprintf("%x", hs)
}

// CheckBySalt
// check HmacHashWithSalt
func CheckBySalt(check, hash, salt string) bool {
	expectedMAC := HmacHashWithSalt(check, salt)
	return hmac.Equal([]byte(hash), []byte(expectedMAC))
}
