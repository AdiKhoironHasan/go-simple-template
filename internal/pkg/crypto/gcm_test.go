package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGCM(t *testing.T) {
	secretKey := "very secret"
	gcm, err := New(secretKey)
	assert.NoError(t, err)

	plaintext := "hello world"

	t.Logf("plaintext: %s", plaintext)

	ciphertext, err := gcm.Encrypt(plaintext)
	assert.NoError(t, err)

	t.Logf("ciphertext: %s", ciphertext)

	decrypted, err := gcm.Decrypt(ciphertext)
	assert.NoError(t, err)
	assert.Equal(t, plaintext, decrypted)
}
