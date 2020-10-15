package hash

import (
	"github.com/l552121229/golang-tools/hash/fnv64a"
)

// GetFnv64a returns a new 64-bit FNV-1a Hasher which makes no memory allocations.
// Its Sum64 method will lay the value out in big-endian byte order.
// See https://en.wikipedia.org/wiki/Fowler–Noll–Vo_hash_function
func GetFnv64a() fnv64a.Hash {
	return fnv64a.Hash{}
}
