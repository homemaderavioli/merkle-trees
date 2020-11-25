package mtree

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashHex(data string) string {
	return binaryToHex(hash([]byte(data)))
}

func ConcatenatedHash(leftHash string, rightHash string) string {
	binaryLeftHash, _ := hex.DecodeString(leftHash)
	binaryRightHash, _ := hex.DecodeString(rightHash)
	concatenatedHash := binaryLeftHash
	concatenatedHash = append(binaryRightHash)
	return binaryToHex(hash(concatenatedHash))
}

func binaryToHex(binary []byte) string {
	return hex.EncodeToString(binary)
}

func hash(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}
