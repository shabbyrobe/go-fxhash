package fxhash

// The original source uses host endian ordering; we favour little endian instead
// by default so the hash is portable.

func leUint64(b []byte) uint64 {
	_ = b[7]
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
}

func leUint32(b []byte) uint64 {
	_ = b[3]
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24
}

func leUint16(b []byte) uint64 {
	_ = b[1]
	return uint64(b[0]) | uint64(b[1])<<8
}

func leUint64Str(s string) uint64 {
	_ = s[7]
	return uint64(s[0]) | uint64(s[1])<<8 | uint64(s[2])<<16 | uint64(s[3])<<24 |
		uint64(s[4])<<32 | uint64(s[5])<<40 | uint64(s[6])<<48 | uint64(s[7])<<56
}

func leUint32Str(s string) uint64 {
	_ = s[3]
	return uint64(s[0]) | uint64(s[1])<<8 | uint64(s[2])<<16 | uint64(s[3])<<24
}

func leUint16Str(s string) uint64 {
	_ = s[1]
	return uint64(s[0]) | uint64(s[1])<<8
}
