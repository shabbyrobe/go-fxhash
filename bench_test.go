package fxhash

import (
	"fmt"
	"math/rand"
	"testing"
)

var hashSizes = []int{1, 2, 4, 8, 16, 32, 128}

var BenchHash64Result uint64

func BenchmarkFNV1a64(b *testing.B) {
	data := make([]byte, 128)
	rand.Read(data)
	for idx, sz := range hashSizes {
		b.Run(fmt.Sprintf("%d-%d", idx, sz), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				BenchHash64Result = fnv1a64(data[:sz])
			}
		})
	}
}

func BenchmarkFxHash(b *testing.B) {
	data := make([]byte, 128)
	rand.Read(data)
	for idx, sz := range hashSizes {
		b.Run(fmt.Sprintf("%d-%d", idx, sz), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				BenchHash64Result = Sum64(data[:sz])
			}
		})
	}
}
