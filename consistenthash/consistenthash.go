package consistenthash

import (
	"sync"

	"github.com/spaolacci/murmur3"
)

const (
	maxWeight       = 100
	minReplicaCount = 100

	prime = 359445410222021406858304963543 // random prime number
)

type HashFn func([]byte) uint64

type ConsistentHash struct {
	hashFn HashFn

	replicaCount uint

	mut        sync.RWMutex
	keys       []uint64
	ring       map[uint64][]any
	nodeLookup map[uint64]any
}

func New(opts ...Option) *ConsistentHash {
	ch := &ConsistentHash{
		hashFn:     murmur3.Sum64,
		keys:       make([]uint64, 0),
		ring:       make(map[uint64][]any),
		nodeLookup: make(map[uint64]any),
	}

	for _, opt := range opts {
		opt(ch)
	}

	return ch
}
