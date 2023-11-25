package consistenthash

type Option func(*ConsistentHash)

func SetReplicas(count uint64) Option {
	return func(ch *ConsistentHash) {
		if count < minReplicaCount {
			count = minReplicaCount
		}
		ch.replicaCount = uint(count)
	}
}

func SetHashFn(fn HashFn) Option {
	return func(ch *ConsistentHash) {
		ch.hashFn = fn
	}
}
