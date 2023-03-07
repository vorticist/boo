package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"gitlab.com/vorticist/logger"
	"io"
)

var c *crypt

func New(key string) Crypt {
	if c == nil {
		c = &crypt{
			key: key,
		}
	}
	return c
}

func Get() Crypt {
	return c
}

type Crypt interface {
	Encrypt(string) (string, error)
	Decrypt(string) (string, error)
}

type crypt struct {
	key string
}

func (c *crypt) Encrypt(text string) (string, error) {
	key, err := hex.DecodeString(c.key)
	if err != nil {
		logger.Errorf("failed to decode key: %v", err)
		return "", err
	}
	plaintext := []byte(text)
	block, err := aes.NewCipher(key)
	if err != nil {
		logger.Errorf("failed to create cipher: %v", err)
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		logger.Errorf("failed to create aesGCM: %v", err)
		return "", err
	}
	nonce := make([]byte, aesGCM.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		logger.Errorf("failed to read data: %v", err)
		return "", err
	}
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext), nil
}

func (c *crypt) Decrypt(text string) (string, error) {
	key, err := hex.DecodeString(c.key)
	if err != nil {
		logger.Errorf("failed to decode key: %v", err)
		return "", err
	}

	enc, err := hex.DecodeString(text)
	if err != nil {
		logger.Errorf("failed to decode encrypted: %v", err)
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		logger.Errorf("failed to create cipher: %v", err)
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		logger.Errorf("failed to create aesGCM: %v", err)
		return "", err
	}
	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		logger.Errorf("failed to open cipher: %v", err)
		return "", err
	}

	return fmt.Sprintf("%s", plaintext), nil
}
