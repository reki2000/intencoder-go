package fistel

import (
	"crypto/sha256"
	"encoding/base32"
	"encoding/binary"
	"math/rand"
)

type Scrambler struct {
	salt    []byte
	encoder *base32.Encoding
}

func NewScrambler(salt string) *Scrambler {
	return &Scrambler{
		salt:    []byte(salt),
		encoder: base32.StdEncoding.WithPadding(base32.NoPadding),
	}
}

func (f *Scrambler) Encode(data int64) string {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, uint64(data))

	nonce := byte(rand.Int31())

	key := sha256.Sum256(append(f.salt, nonce))

	i := 7
	for bytes[i] == 0 && i > 3 {
		i--
	}

	value := make([]byte, i+1)
	for i := 0; i < len(value); i++ {
		value[i] = bytes[i] ^ key[i]
	}

	value = append(value, nonce)

	return f.encoder.EncodeToString(value)
}

func (f *Scrambler) Decode(data string) (int64, error) {
	bytes, err := f.encoder.DecodeString(data)
	if err != nil {
		return 0, err
	}

	nonce := bytes[len(bytes)-1]
	key := sha256.Sum256(append(f.salt, nonce))

	bytes8 := make([]byte, 8)
	for i := 0; i < len(bytes)-1; i++ {
		bytes8[i] = bytes[i] ^ key[i]
	}

	decrypted := binary.LittleEndian.Uint64(bytes8)
	return int64(decrypted), nil
}
