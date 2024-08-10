package fistel

import (
	"crypto/sha512"
	"encoding/binary"
	"fmt"
)

type Fistel32 struct {
	keys [32]uint32
}

func NewFistel32(salt string) *Fistel32 {
	// convert salt to sha512
	bytes := sha512.Sum512([]byte(salt))

	// split into uint32 slices (uses only uint16 size)
	f := Fistel32{}
	for i := 0; i < 32; i++ {
		f.keys[i] = uint32(binary.BigEndian.Uint16(bytes[i*2 : i*2+2]))
	}

	fmt.Printf("fistel32: %+v\n", f)

	return &f
}

func round32(data uint32, key uint32) uint32 {
	left := data >> 16
	right := data & 0xFFFF
	newRight := left ^ (right ^ key)
	return (right << 16) | newRight
}

func (f *Fistel32) Encrypt(data uint32) uint32 {
	x := data
	for i := 0; i < len(f.keys); i++ {
		x = round32(x, f.keys[i])
	}
	return x
}

func (f *Fistel32) Decrypt(data uint32) uint32 {
	x := data
	for i := len(f.keys) - 1; i >= 0; i-- {
		x = round32(x, f.keys[i])
	}
	return x
}
