package consistenthash

const (
	maxWeight   = 100
	minReplicas = 100

	prime = 359445410222021406858304963543
)

type HashFn func([]byte) []byte

type ConsistentHash struct {
	hashFn HashFn

	replicaCount uint64
	weight       uint32
	ring         []uint64
	nodeLookup   map[uint64]any
}
