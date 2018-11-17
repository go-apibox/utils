package utils

import (
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/go-apibox/types"
)

var allLetters = []rune("023456789abcdefghjkmnopqrstuvwxyzABCDEFGHJKMNOPQRSTUVWXYZ")
var alphaLetters = []rune("abcdefghjkmnopqrstuvwxyzABCDEFGHJKMNOPQRSTUVWXYZ")
var digitLetters = []rune("023456789")

func randStringN(letters []rune, n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}

func RandStringN(n int) string {
	return randStringN(allLetters, n)
}

func RandString() string {
	return RandStringN(16)
}

func RandAlphaStringN(n int) string {
	return randStringN(alphaLetters, n)
}

func RandAlphaString() string {
	return RandAlphaStringN(16)
}

func RandDigitStringN(n int) string {
	return randStringN(digitLetters, n)
}

func RandDigitString() string {
	return RandAlphaStringN(16)
}

func RandUint() uint {
	if types.MaxInt == math.MaxInt32 {
		return uint(RandUint32())
	} else {
		return uint(RandUint64())
	}
}

func RandUintIn(low, hi uint) uint {
	if types.MaxInt == math.MaxInt32 {
		return uint(RandUint32In(uint32(low), uint32(hi)))
	} else {
		return uint(RandUint64In(uint64(low), uint64(hi)))
	}
}

func RandUint32() uint32 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Uint32()
}

func RandUint32In(low, hi uint32) uint32 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return low + uint32(r.Int31n(int32(hi-low)))
}

func RandUint64() uint64 {
	// javascript 整数最大只能支持到：9007199254740992
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return uint64(r.Int63n(9007199254740993))
}

func RandUint64In(low, hi uint64) uint64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return low + uint64(r.Int63n(int64(hi-low)))
}

func RandDatePrefixUint64() uint64 {
	t := time.Now().Format("20060102")
	datePart, err := strconv.ParseUint(t, 10, 64)
	if err != nil {
		return 0
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randPart := uint64(r.Int31n(100000000))
	return datePart*100000000 + randPart
}
