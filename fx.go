package fxhash

import (
	"math/bits"
)

func Sum64(b []byte) uint64 {
	return Append64(initial, b)
}

func Append64(hash uint64, b []byte) uint64 {
	for len(b) >= 8 {
		hash = (bits.RotateLeft64(hash, rotate) ^ leUint64(b)) * seed
		b = b[8:]
	}

	if len(b) >= 4 {
		hash = (bits.RotateLeft64(hash, rotate) ^ leUint32(b)) * seed
		b = b[4:]
	}

	if len(b) >= 2 {
		hash = (bits.RotateLeft64(hash, rotate) ^ leUint16(b)) * seed
		b = b[2:]
	}

	if len(b) >= 1 {
		hash = (bits.RotateLeft64(hash, rotate) ^ uint64(b[0])) * seed
	}

	return hash
}

func Sum64String(s string) uint64 {
	return Append64String(initial, s)
}

func Append64String(hash uint64, s string) uint64 {
	for len(s) >= 8 {
		hash = (bits.RotateLeft64(hash, rotate) ^ leUint64Str(s)) * seed
		s = s[8:]
	}

	if len(s) >= 4 {
		hash = (bits.RotateLeft64(hash, rotate) ^ leUint32Str(s)) * seed
		s = s[4:]
	}

	if len(s) >= 2 {
		hash = (bits.RotateLeft64(hash, rotate) ^ leUint16Str(s)) * seed
		s = s[2:]
	}

	if len(s) >= 1 {
		hash = (bits.RotateLeft64(hash, rotate) ^ uint64(s[0])) * seed
	}

	return hash
}
