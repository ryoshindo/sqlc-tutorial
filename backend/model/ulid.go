package model

import (
	"crypto/rand"
	"sync"

	"github.com/oklog/ulid/v2"
)

var (
	defaultEntropySource *ulid.MonotonicEntropy
	entroypySourceMtx    sync.Mutex
)

func init() {
	defaultEntropySource = ulid.Monotonic(rand.Reader, 0)
}

func NewUlid() ulid.ULID {
	entroypySourceMtx.Lock()
	defer entroypySourceMtx.Unlock()

	return ulid.MustNew(ulid.Now(), defaultEntropySource)
}

func NewUlidString() string {
	return NewUlid().String()
}
