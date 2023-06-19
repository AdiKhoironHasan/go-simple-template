package cryptox

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
)

func generateKeyPair() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)

}

func pemEncodePrivateKey(privateKey *rsa.PrivateKey) []byte {
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	return pem.EncodeToMemory(block)
}

func pemEncodePublicKey(publicKey *rsa.PublicKey) []byte {
	bytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		panic(err)
	}

	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: bytes,
	}

	return pem.EncodeToMemory(block)
}

func encryptWithPublicKey(message []byte, publicKey *rsa.PublicKey) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, publicKey, message)
}

func decryptWithPrivateKey(ciphertext []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
}

func HashSHA256(message string) string {
	hash := sha256.Sum256([]byte(message))
	hashString := hex.EncodeToString(hash[:])

	return hashString
}
