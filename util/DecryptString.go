package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
)

func Decrypt(data []byte, passphrase string) []byte {
	plaintext := data
	defer func() {
		if err := recover(); err != nil {
			plaintext = data
		}
	}()
	hasher := md5.New()
	hasher.Write([]byte(passphrase))
	key := []byte(hex.EncodeToString(hasher.Sum(nil)))
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte(data)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return []byte(data)
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err = gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return []byte(data)
	}
	return plaintext
}
