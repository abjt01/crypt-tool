package filecrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"io"
	"os"

	"golang.org/x/crypto/pbkdf2"
)

func Encrypt(source string, password []byte) {
	plaintext, err := os.ReadFile(source)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}

	dk := pbkdf2.Key(password, nonce, 4096, 32, sha1.New)

	block, err := aes.NewCipher(dk)
	if err != nil {
		panic(err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	ciphertext = append(ciphertext, nonce...)

	if err := os.WriteFile(source, ciphertext, 0644); err != nil {
		panic(err)
	}
}

func Decrypt(source string, password []byte) {
	ciphertext, err := os.ReadFile(source)
	if err != nil {
		panic(err)
	}

	nonce := ciphertext[len(ciphertext)-12:]
	cipherBody := ciphertext[:len(ciphertext)-12]

	dk := pbkdf2.Key(password, nonce, 4096, 32, sha1.New)

	block, err := aes.NewCipher(dk)
	if err != nil {
		panic(err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	plaintext, err := aesgcm.Open(nil, nonce, cipherBody, nil)
	if err != nil {
		panic("Decryption failed — wrong password or corrupted file")
	}

	if err := os.WriteFile(source, plaintext, 0644); err != nil {
		panic(err)
	}
}
