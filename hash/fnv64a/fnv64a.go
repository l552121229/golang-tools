package fnv64a

// Hash Hash
type Hash struct{}

const (
	// offset64 FNVa offset basis. See https://en.wikipedia.org/wiki/Fowler–Noll–Vo_hash_function#FNV-1a_hash
	offset64 = 14695981039346656037
	// prime64 FNVa prime value. See https://en.wikipedia.org/wiki/Fowler–Noll–Vo_hash_function#FNV-1a_hash
	prime64 = 1099511628211
)

// Sum64 gets the string and returns its uint64 hash value.
func (f Hash) Sum64(key string) uint64 {
	var hash uint64 = offset64
	for _, k := range key {
		hash = (hash ^ uint64(k)) * prime64
	}

	return hash
}
