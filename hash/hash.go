package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"time"
)

func hashMD5(data []byte) []byte {
	hash := md5.Sum(data)
	return hash[:]
}

func hashSHA1(data []byte) []byte {
	hash := sha1.Sum(data)
	return hash[:]
}

func hashSHA256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

func hashSHA512(data []byte) []byte {
	hash := sha512.Sum512(data)
	return hash[:]
}

func benchmarkHashing(fn func([]byte) []byte, data []byte) time.Duration {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		fn(data)
	}
	return time.Since(start)
}

func main() {
	data := []byte("hello world")
	fmt.Printf("Hashing 'hello world' with different algorithms:\n")

	md5Time := benchmarkHashing(hashMD5, data)
	sha1Time := benchmarkHashing(hashSHA1, data)
	sha256Time := benchmarkHashing(hashSHA256, data)
	sha512Time := benchmarkHashing(hashSHA512, data)

	fmt.Printf("MD5   : %v\n", md5Time)
	fmt.Printf("SHA1  : %v\n", sha1Time)
	fmt.Printf("SHA256: %v\n", sha256Time)
	fmt.Printf("SHA512: %v\n", sha512Time)
}
