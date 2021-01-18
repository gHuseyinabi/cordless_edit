package util

import (
	b64 "encoding/base64"
)

func DecryptBase64(data string) []byte {
	context, _ := b64.StdEncoding.DecodeString(data)
	return context
}

func EncryptBase64(data []byte) string {
	context := b64.StdEncoding.EncodeToString(data)
	return context
}
