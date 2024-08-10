package main

import (
	"fmt"

	"github.com/reki2000/intencoder-go"
)

func demo(f *intencoder.IntEncoder, value uint64) {
	encoded := f.Encode(value)
	decoded, err := f.Decode(encoded)
	if err != nil {
		fmt.Printf("%s %s\n", encoded, err)
		return
	}
	fmt.Printf("%s %d 0x%x\n", encoded, decoded, decoded)

}
func main() {
	e := intencoder.NewIntEncoder("intencoder-go").WithDelimiter(":").WithMinByteLength(5)
	demo(e, 0)
	demo(e, 0x11)
	demo(e, 0x1122)
	demo(e, 0x112233)
	demo(e, 0x11223344)
	demo(e, 0x1122334455)
	demo(e, 0x112233445566)
	demo(e, 0x11223344556677)
	demo(e, 0x1122334455667788)
}
