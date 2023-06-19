package cryptox

import (
	"crypto/rsa"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeDecode(t *testing.T) {
	message := []byte("This message is very secret.")

	privateKey, err := generateKeyPair()
	if err != nil {
		panic(err)
	}

	privateKeyPEM := pemEncodePrivateKey(privateKey)
	fmt.Printf("Private key:\n%s\n", privateKeyPEM)

	publicKey := privateKey.Public().(*rsa.PublicKey)
	publicKeyPEM := pemEncodePublicKey(publicKey)
	fmt.Printf("Public key:\n%s\n", publicKeyPEM)

	ciphertext, err := encryptWithPublicKey(message, publicKey)
	assert.NoError(t, err)

	fmt.Printf("Ciphertext: %x\n", ciphertext)
	plaintext, err := decryptWithPrivateKey(ciphertext, privateKey)
	assert.NoError(t, err)

	fmt.Printf("Plaintext: %s\n", plaintext)

	assert.Equal(t, message, plaintext)
}

func TestHashing(t *testing.T) {
	message := "Hello, World!"

	hashString := HashSHA256(message)

	fmt.Println("Plain Text : ", message)
	fmt.Println("SHA256 Hash : ", hashString)
}
