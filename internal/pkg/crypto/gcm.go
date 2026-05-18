package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
)

// Service is a struct that contains the GCM instance
type Service struct {
	gcm cipher.AEAD
}

// New creates a new instance of the crypto service
func New(secretKey string) (*Service, error) {
	if secretKey == "" {
		return nil, errors.New("secret key is empty")
	}

	// Hash key input
	hasher := sha256.New()
	hasher.Write([]byte(secretKey))
	keyBytes := hasher.Sum(nil)

	// Create AES cipher block
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES cipher block: %w", err)
	}

	// Wrap with GCM (Galois/Counter Mode)
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat GCM: %w", err)
	}

	return &Service{
		gcm: aesGCM,
	}, nil
}

// Encrypt encrypts a plaintext string to a ciphertext string (Base64).
func (s *Service) Encrypt(plaintext string) (string, error) {
	if plaintext == "" {
		return "", errors.New("plaintext kosong")
	}

	// Create Nonce (Number used once)
	nonce := make([]byte, s.gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("failed to create nonce: %w", err)
	}

	// Encrypt (Seal)
	// Result: nonce + ciphertext + tag (integrity tag is created automatically by GCM)
	ciphertext := s.gcm.Seal(nonce, nonce, []byte(plaintext), nil)

	// Encode to Base64
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt decrypts a ciphertext string (Base64) back to a plaintext string.
func (s *Service) Decrypt(cryptoText string) (string, error) {
	if cryptoText == "" {
		return "", errors.New("ciphertext is empty")
	}

	// Decode Base64
	data, err := base64.StdEncoding.DecodeString(cryptoText)
	if err != nil {
		return "", fmt.Errorf("invalid base64 format: %w", err)
	}

	nonceSize := s.gcm.NonceSize()
	if len(data) < nonceSize {
		return "", errors.New("ciphertext is corrupted or too short")
	}

	// Separate Nonce and Ciphertext
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	// Decrypt (Open)
	plaintext, err := s.gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("failed to decrypt (wrong key or manipulated data): %w", err)
	}

	return string(plaintext), nil
}
