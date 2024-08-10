# intencoder-go

A UInt64 obfuscation encoder for sequencial ID

## Motivation

In a simple RDB-based web system, using sequential IDs for entities is not ideal. Such practice allows for the estimation of the business size and increases the risk of ID-guessing attacks.

This library provides a reversible transformation to obfuscate a uint64 into a hard-to-guess string. The resulting string ranges from a minimum of 7 to a maximum of 15 characters (or an array with a minimum of 4 bytes to a maximum of 9 bytes).

To enhance human readability and obfuscate the ID further, the encoder utilizes a restricted character set based on base32 encoding. This avoids characters like "0" and "O", "1" and "I", and especially "Q" and "9", as these pairs are easily confused in pronunciation in Japanese.

- up to 2^24-1 : 7
- up to 2^32-1 : 8
- up to 2^40-1 : 10
- up to 2^48-1 : 12
- up to 2^56-1 : 13
- up to 2^64-1 : 15

Example:

| Code              | Decimal Value       | Hexadecimal Value |
|-------------------|---------------------|-------------------|
| KM4F-DLY          | 0                   | 0x0               |
| DWE3-QNI          | 17                  | 0x11              |
| Z3FR-Y4Q          | 4386                | 0x1122            |
| GBDK-3DY          | 1122867             | 0x112233          |
| WJF3-LN4L         | 287454020           | 0x11223344        |
| A6NH-TF27-HI      | 73588229205         | 0x1122334455      |
| 5EZH-4LSJ-TKRQ    | 18838586676582      | 0x112233445566    |
| W43I-TY3I-MLSI-C  | 4822678189205111    | 0x11223344556677  |
| VXQB-FS37-LKNN-HAA| 1234605616436508552 | 0x1122334455667788|

#### Learning

Initially, I attempted to use the Feistel algorithm, but it did not effectively obscure the change in bits across a sequence of original values.

To address this, I added a nonce byte and created a different SHA-256 hash from a salt and nonce. The encryption is performed using a simple XOR operation with the hash.

#### other notes
- golang hash function performance comparison:
```
MD5   : 1.074252ms
SHA1  : 1.11365ms
SHA256: 771.172µs
SHA512: 2.29784ms
```
