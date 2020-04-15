package fxhash

import (
	"math/bits"
)

func New() *Digest {
	return Seed(initial)
}

func Seed(v uint64) *Digest {
	var d Digest
	d.seed = v
	d.Reset()
	return &d
}

type Digest struct {
	seed  uint64
	state uint64
	left  [8]byte
	leftn int
}

func (d *Digest) Size() int      { return 8 }
func (d *Digest) BlockSize() int { return 8 }

func (d *Digest) Sum64() uint64 {
	if d.leftn == 0 {
		return d.state
	}
	return Append64(d.state, d.left[:d.leftn])
}

func (d *Digest) Write(b []byte) (n int, err error) {
	n = len(b)
	for len(b) >= 8 {
		d.state = (bits.RotateLeft64(d.state, rotate) ^ leUint64(b)) * seed
		b = b[8:]
	}
	if len(b) > 0 {
		d.leftn = copy(d.left[:], b)
	}
	return n, nil
}

func (d *Digest) Sum(b []byte) []byte {
	sum := d.Sum64()
	return append(b,
		byte(sum>>56), byte(sum>>48), byte(sum>>40), byte(sum>>32),
		byte(sum>>24), byte(sum>>16), byte(sum>>8), byte(sum))
}

func (d *Digest) Reset() {
	d.state = d.seed
	d.leftn = 0
	d.left = [8]byte{}
}
