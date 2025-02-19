package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var currRand *rand.Rand

func init() {
	currRand = rand.New(rand.NewSource(time.Now().UnixNano()))
}
 
// RandomInt generates a randm integer between min and max
func RandomInt(min, max int64) int64 {
	return min + currRand.Int63n(max - min + 1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder

	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[currRand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 100)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	currencies := [3]string{EUR, USD, CAD}

	n := len(currencies)

	return currencies[currRand.Intn(n)]
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
