package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
)

type AesCipher struct {
	key []byte
}

func NewAesCipher(key string) (*AesCipher, error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, errors.New("invalid key size")
	}
	return &AesCipher{[]byte(key)}, nil
}

func (c *AesCipher) Encrypt(plaintext string) (string, error) {
	block, err := aes.NewCipher(c.key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	plaintextBytes := PKCS5Padding([]byte(plaintext), blockSize)

	iv := make([]byte, blockSize)
	if _, err = rand.Read(iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	ciphertextBytes := make([]byte, len(plaintextBytes))
	mode.CryptBlocks(ciphertextBytes, plaintextBytes)

	resultBytes := make([]byte, len(ciphertextBytes)+blockSize)
	copy(resultBytes[:blockSize], iv)
	copy(resultBytes[blockSize:], ciphertextBytes)
	return base64.StdEncoding.EncodeToString(resultBytes), nil
}

func (c *AesCipher) Decrypt(ciphertext string) (string, error) {
	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	blockSize := aes.BlockSize
	if len(ciphertextBytes) < 2*blockSize {
		return "", errors.New("ciphertext too short")
	}

	iv := ciphertextBytes[:blockSize]
	ciphertextBytes = ciphertextBytes[blockSize:]

	block, err := aes.NewCipher(c.key)
	if err != nil {
		return "", err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	plaintextBytes := make([]byte, len(ciphertextBytes))
	mode.CryptBlocks(plaintextBytes, ciphertextBytes)

	plaintextBytes = PKCS5UnPadding(plaintextBytes)
	return string(plaintextBytes), nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

func PKCS5UnPadding(plaintext []byte) []byte {
	length := len(plaintext)
	unPadding := int(plaintext[length-1])
	return plaintext[:(length - unPadding)]
}
