package common

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"study-service/config"
)

//var iv = []byte{34, 45, 55, 22, 87, 31, 29, 76, 84, 32, 88, 93, 62, 36, 15, 07}
var iv = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func decodeBase64(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func GetDecrypt(text string) string {
	key := config.Config.Encrypt.EncryptKey
	return Decrypt(key, text)
}

func SetEncrypt(text string) string {
	key := config.Config.Encrypt.EncryptKey
	return Encrypt(key, text)
}

func Encrypt(key, text string) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	plaintext := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	return encodeBase64(ciphertext)
}

func Decrypt(key, text string) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	ciphertext := decodeBase64(text)
	cfb := cipher.NewCFBEncrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	cfb.XORKeyStream(plaintext, ciphertext)
	return string(plaintext)
}

func EncryptSha256(text string) string {
	hash := sha256.New()
	hash.Write([]byte(text))

	return hex.EncodeToString(hash.Sum(nil))
}
