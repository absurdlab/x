package codecx

import "crypto/sha256"

// SHA256 hashes data in SHA-256.
func SHA256(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	return h.Sum(nil)
}

// SHA256ToString hashes data in SHA-256 and then runs it through the given Encoder.
func SHA256ToString(data []byte, encoder Encoder) string {
	return encoder(SHA256(data))
}
