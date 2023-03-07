package crypt

import (
	"github.com/vorticist/boo/client"
	"gitlab.com/vorticist/logger"
	"testing"
)

func TestCrypt_Encrypt(t *testing.T) {
	c := client.New()
	keyStr, err := c.GetCryptKey()
	if err != nil {
		t.Errorf("failed to get crypt key: %v", err)
		return
	}

	cr := New(keyStr)

	msg := "hello world"
	encrypted, err := cr.Encrypt(msg)
	if err != nil {
		t.Errorf("failed to encrypt message: %v", err)
		return
	}

	logger.Infof("encrypted message: %v", encrypted)
}

func TestCrypt_Decrypt(t *testing.T) {
	c := client.New()
	keyStr, err := c.GetCryptKey()
	if err != nil {
		t.Errorf("failed to get crypt key: %v", err)
		return
	}

	cr := New(keyStr)
	encrypted := "364272aa3ca60e482357b8f736e37782df15d71d87dd3b7827a3efe80fc9e443650825dea19261"

	decrypted, err := cr.Decrypt(encrypted)
	if err != nil {
		t.Errorf("failed to decrypt: %v", err)
		return
	}

	if decrypted != "hello world" {
		t.Errorf("invalid decrypted message: %v", decrypted)
		return
	}

	logger.Infof("decrypted: %v", decrypted)
}
