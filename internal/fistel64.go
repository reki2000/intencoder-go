package fistel

import (
	"crypto/sha512"
	"encoding/binary"
	"fmt"
)

type Fistel64 struct {
	keys [16]uint32
}

func NewFistel64(salt string) *Fistel64 {
	// convert salt to sha512
	bytes := sha512.Sum512([]byte(salt))

	// split into uint32 slices
	f := Fistel64{}
	for i := 0; i < 16; i++ {
		f.keys[i] = binary.BigEndian.Uint32(bytes[i*4 : i*4+4])
	}

	fmt.Printf("fistel64: %+v\n", f)

	return &f
}

func round64(data uint64, key uint32) uint64 {
	left := data >> 32
	right := data & 0xFFFFFFFF
	newRight := left ^ (right ^ uint64(key))
	return (right << 32) | newRight
}

func (f *Fistel64) Encrypt(data uint64) uint64 {
	x := data
	for i := 0; i < len(f.keys); i++ {
		x = round64(x, f.keys[i])
	}
	return x
}

func (f *Fistel64) Decrypt(data uint64) uint64 {
	x := data
	for i := len(f.keys) - 1; i >= 0; i-- {
		x = round64(x, f.keys[i])
	}
	return x
}
