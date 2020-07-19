package char

import (
	"math/rand"
	"time"
)

const (
	charStart = 33  // !
	charEnd   = 126 // ~
	charLen   = charEnd - charStart
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandKey(length int) string {
	var result = make([]byte, length, length)

	for i := 0; i < length; i++ {
		result[i] = byte(r.Uint32()%charLen + charStart)
	}

	return string(result)
}
