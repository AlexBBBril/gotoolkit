package random

import (
	"math/rand"
	"strconv"
	"time"
	"unsafe"
)

// Int generating a pseudo random integer between min and max
func Int(min, max int) int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max-min) + min
}

// IntAsString generating a pseudo random integer, in string format, in the range between min and max.
func IntAsString(min, max int) string {
	return strconv.Itoa(Int(min, max))
}

// String character n length generation of a pseudo random string
func String(n int) string {
	const (
		letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

		letterIdxBits = 6
		letterIdxMask = 1<<letterIdxBits - 1
		letterIdxMax  = 63 / letterIdxBits
	)

	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)

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

	return *(*string)(unsafe.Pointer(&b))
}
