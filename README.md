# Motivation

In a simple RDB-based web system, using sequential IDs for entities is not ideal. Such practice allows for the estimation of the business size and increases the risk of ID-guessing attacks.

This library offers a reversible transformation that converts a uint64 into a hard-to-guess string ranging from a minimum of 8 to a maximum of 16 characters (or an array with a minimum of 5 bytes to a maximum of 9 bytes).

To enhance human readability and because the original sequence usually falls within the uint32 range, the encoder generates a shorter string for smaller values:

- up to 2^32-1 : 8
- up to 2^40-1 : 10
- up to 2^48-1 : 12
- up to 2^56-1 : 14
- up to 2^64-1 : 16

Example:

| original (dec) | original (hex) | encoded | 
|-----------------------|-------------------|-------------------| 
| 0 | 0x0 | DMLLDX3O |
| 1234567890 | 0x499602d2 | GXG5B2IL | 
| 1234567891 | 0x499602d3 | 4BA474BD |
| 1234567890123456789 | 0x112210f47de98115| 6Y4GGMCXBIYGY5A |
| 1234567890123456790 | 0x112210f47de98116| QPQYW3UBBCIGDEQ | 
| 73588229205 | 0x1122334455 | HRTNUHBJUU |
| 18838586676582 | 0x112233445566 | QGNAFE2254FQ |
| 4822678189205111 | 0x11223344556677 | 6PQW2EHIQI6ZM |
| 1234605616436508552 | 0x1122334455667788| 6HE2DKT2LLZ2LGA |

# Learning

Initially, I attempted to use the Feistel algorithm, but it did not effectively obscure the change in bits across a sequence of original values.

To address this, I added a nonce byte and created a different SHA-256 hash from a salt and nonce. The encryption is performed using a simple XOR operation.

Feel free to ask if you need more details on how to implement this, or if you have any specific questions about the transformation or the encoding process!