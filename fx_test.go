package fxhash

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"testing"
)

func TestSum64(t *testing.T) {
	for idx, tc := range []struct {
		h  uint64
		in []byte
	}{
		{0x5d84ea10a456a1f7, []byte{}},

		// Problematic property of this hash: leading zeros
		{0xedf92e244033ccc7, []byte{0}},
		{0xedf92e244033ccc7, []byte{0, 0}},
		{0x02852fbc6159ed41, []byte{0, 0, 0}},
		{0xedf92e244033ccc7, []byte{0, 0, 0, 0}},
		{0x02852fbc6159ed41, []byte{0, 0, 0, 0, 0}},
		{0x02852fbc6159ed41, []byte{0, 0, 0, 0, 0, 0}},

		{0x022a0ce3e1b3ea11, []byte("f")},
		{0x448aa4cb687cf911, []byte("fo")},
		{0x1715e832ad63b953, []byte("foo")},
		{0x9a06525b15eff911, []byte("foob")},
		{0x66b8309f99ba2dba, []byte("fooba")},
		{0xd18719a7b7a0f3ba, []byte("foobar")},
		{0xef3eebbaba9a4e02, []byte("foobar ")},
		{0x25cddad015eff911, []byte("foobar b")},
		{0xf8ddd42fd9bbe162, []byte("foobar bazqux")},
	} {
		t.Run(fmt.Sprintf("%d", idx), func(t *testing.T) {
			result := Sum64(tc.in)
			if result != tc.h {
				t.Fatalf("0x%016x != 0x%016x", result, tc.h)
			}

			stringCmp := Sum64String(string(tc.in))
			if result != stringCmp {
				t.Fatal()
			}

			digest := New()
			n, err := digest.Write(tc.in)
			if err != nil {
				t.Fatal(err)
			}
			if n != len(tc.in) {
				t.Fatal()
			}
			if result != digest.Sum64() {
				t.Fatalf("0x%016x != 0x%016x", result, digest.Sum64())
			}
		})
	}
}

func TestSum64Distribution(t *testing.T) {
	const iter = 100000

	rng := rand.NewSource(0).(rand.Source64)

	num := make([]byte, 8)
	bits := make([]int, 64)

	for i := 0; i < iter; i++ {
		n := rng.Uint64()
		binary.LittleEndian.PutUint64(num, n)
		nHash := Sum64(num)

		for i := 0; i < 64; i++ {
			bit := (uint64(1) << i)
			if nHash&bit != 0 {
				bits[i]++
			}
		}
	}

	printResult("fx dist", bits, iter)
}

func TestSum64Avalanche(t *testing.T) {
	const iter = 100000

	rng := rand.NewSource(0).(rand.Source64)

	num := make([]byte, 8)
	flips := make([]int, 64)
	tests := 0

	for i := 0; i < iter; i++ {
		n := rng.Uint64()
		binary.LittleEndian.PutUint64(num, n)
		nHash := Sum64(num)

		for i := 0; i < 64; i++ {
			tests++
			bit := (uint64(1) << i)
			v := n ^ bit

			binary.LittleEndian.PutUint64(num, v)
			vHash := Sum64(num)

			for j := 0; j < 64; j++ {
				if ((nHash^vHash)>>j)&1 != 0 {
					flips[j]++
				}
			}
		}
	}

	printResult("fx avalanche", flips, tests)
}

func TestFNV1aAvalanche(t *testing.T) {
	const iter = 100000

	rng := rand.NewSource(0).(rand.Source64)

	num := make([]byte, 8)
	flips := make([]int, 64)
	tests := 0

	for i := 0; i < iter; i++ {
		n := rng.Uint64()
		binary.LittleEndian.PutUint64(num, n)
		nHash := fnv1a64(num)

		for i := 0; i < 64; i++ {
			tests++
			bit := (uint64(1) << i)
			v := n ^ bit

			binary.LittleEndian.PutUint64(num, v)
			vHash := fnv1a64(num)

			for j := 0; j < 64; j++ {
				if ((nHash^vHash)>>j)&1 != 0 {
					flips[j]++
				}
			}
		}
	}

	printResult("fnv1a avalanche", flips, tests)
}

func printResult(fn string, flips []int, tests int) {
	fmt.Printf("%s:\n", fn)
	cell := 0
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			fmt.Printf(" %0.3f  ", float64(flips[cell])/float64(tests))
			cell++
		}
		fmt.Println()
	}
	fmt.Println()
}

func fnv1a64(b []byte) uint64 {
	const offset64 = 14695981039346656037
	const prime64 = 1099511628211

	var hash uint64 = offset64
	for _, c := range b {
		hash ^= uint64(c)
		hash *= prime64
	}
	return hash
}
