package utility

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

const key string = "the-key-has-to-be-32-bytes-long!"

func EncryptUserPassword(password []byte) ([]byte, error) {
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, password, nil), nil
}

func DecryptUserPassword(password []byte) ([]byte, error) {
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(password) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := password[:nonceSize], password[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}
