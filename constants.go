package fxhash

const (
	rotate = 5
	seed   = uint64(0x51_7c_c1_b7_27_22_0a_95)

	// The Rust implementation "seeds" the hash function with the length
	// of the slice when hashing bytes. We don't emulate this, we provide
	// our own seed.
	//
	// The Rust behaviour can be emulated like so:
	//	var num = make([]byte, 8)
	// 	binary.LittleEndian.PutUint64(num, uint64(len(buf)))
	// 	hash = fxhash.Append64(hash, num)
	// 	hash = fxhash.Append64(hash, buf)
	//
	initial = uint64(0x5d_84_ea_10_a4_56_a1_f7)
)
