package codecx

import (
	"encoding/base64"
	"fmt"
)

// Encoder is a function to turn bytes of raw data to a string
type Encoder func([]byte) string

// Base64Std is an Encoder that encodes raw data in Base64 standard encoding
func Base64Std(raw []byte) string {
	return base64.StdEncoding.EncodeToString(raw)
}

// Hex is an Encoder that encodes raw data in hex format
func Hex(raw []string) string {
	return fmt.Sprintf("%x", raw)
}
