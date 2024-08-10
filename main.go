package main

import (
	"fmt"

	fistel "reki2000.example.com/fistel-go/internal"
)

func demo(f *fistel.Scrambler, value int64) {
	encoded := f.Encode(value)
	decoded, _ := f.Decode(encoded)
	fmt.Printf("%s %d 0x%x\n", encoded, decoded, decoded)

}
func main() {
	e := fistel.NewScrambler("")
	demo(e, 0)
	demo(e, 1234567890)
	demo(e, 1234567891)
	demo(e, 1234567890123456789)
	demo(e, 1234567890123456790)
	demo(e, 0x1122334455)
	demo(e, 0x112233445566)
	demo(e, 0x11223344556677)
	demo(e, 0x1122334455667788)
}
