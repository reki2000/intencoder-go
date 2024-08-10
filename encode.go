package intencoder

import (
	"crypto/sha256"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"strings"
)

type IntEncoder struct {
	salt          []byte
	encoder       *base32.Encoding
	delimiter     string
	minByteLength int
}

func NewIntEncoder(salt string) *IntEncoder {
	return &IntEncoder{
		salt:          []byte(salt),
		encoder:       base32.StdEncoding.WithPadding(base32.NoPadding),
		delimiter:     "-",
		minByteLength: 3,
	}
}

func (encoder *IntEncoder) WithDelimiter(delimiter string) *IntEncoder {
	encoder.delimiter = delimiter
	return encoder
}

func (encoder *IntEncoder) WithMinByteLength(minByteLength int) *IntEncoder {
	encoder.minByteLength = minByteLength
	return encoder
}

func (encoder *IntEncoder) Encode(data uint64) string {
	bytesUint64LE := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytesUint64LE, data)

	nonce := sha256.Sum256(bytesUint64LE)[0]

	key := sha256.Sum256(append(encoder.salt, nonce))

	// shorten the length of the value
	i := 7
	for bytesUint64LE[i] == 0 && i >= encoder.minByteLength {
		i--
	}

	value := make([]byte, i+1)
	for i := 0; i < len(value); i++ {
		value[i] = bytesUint64LE[i] ^ key[i]
	}

	value = append(value, nonce)

	encoded := encoder.encoder.EncodeToString(value)

	// split every 4 characters
	var parts []string
	for i := 0; i < len(encoded); i += 4 {
		end := i + 4
		if end > len(encoded) {
			end = len(encoded)
		}
		parts = append(parts, encoded[i:end])
	}

	return strings.Join(parts, encoder.delimiter)
}

func (encoder *IntEncoder) Decode(data string) (uint64, error) {
	encrypted, err := encoder.encoder.DecodeString(strings.Replace(data, encoder.delimiter, "", -1))
	if err != nil {
		return 0, err
	}

	nonce := encrypted[len(encrypted)-1]
	key := sha256.Sum256(append(encoder.salt, nonce))

	bytesUint64LE := make([]byte, 8)
	for i := 0; i < len(encrypted)-1; i++ {
		bytesUint64LE[i] = encrypted[i] ^ key[i]
	}

	originalNonce := sha256.Sum256(bytesUint64LE)[0]
	if nonce != originalNonce {
		return 0, fmt.Errorf("nonce mismatch")
	}

	decrypted := binary.LittleEndian.Uint64(bytesUint64LE)
	return decrypted, nil
}
