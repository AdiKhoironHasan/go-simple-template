package crypto

import (
	"crypto/rsa"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	message := []byte("{\"user_id\":1,\"exp\":1634025600}")

	privateKey, err := GenerateKeyPair()
	assert.NoError(t, err)

	_ = PemEncodePrivateKey(privateKey)
	t.Log("Private key: [REDACTED]")

	publicKey := privateKey.Public().(*rsa.PublicKey)
	_, err = PemEncodePublicKey(publicKey)
	assert.NoError(t, err)
	t.Log("Public key: [REDACTED]")

	ciphertext, err := EncryptWithPublicKey(message, publicKey)
	assert.NoError(t, err)

	t.Logf("Ciphertext: %x", ciphertext)
	plaintext, err := DecryptWithPrivateKey(ciphertext, privateKey)
	assert.NoError(t, err)

	t.Logf("Plaintext: %s", plaintext)

	assert.Equal(t, message, plaintext)
}

func TestKey(t *testing.T) {
	privateKey, err := GenerateKeyPair()
	require.NoError(t, err)

	privateKey2, err := GenerateKeyPair()
	assert.NoError(t, err)

	privateKeyPEM := PemEncodePrivateKey(privateKey)
	t.Log("Private key: [REDACTED]")

	privateKeyPEM2 := PemEncodePrivateKey(privateKey2)
	t.Log("Private key2: [REDACTED]")

	assert.NotEqual(t, privateKeyPEM, privateKeyPEM2)
}

func TestEncodeDecodeBase64(t *testing.T) {
	message := fmt.Sprintf("key_%d", time.Now().Unix())
	encoded := EncodeBASE64URL(message)
	t.Logf("Encoded: %s", encoded)
	decoded, err := DecodeBASE64(encoded)
	assert.NoError(t, err)
	t.Logf("Decoded: %s", decoded)
	assert.Equal(t, message, decoded)
}

func TestSHA512(t *testing.T) {
	clSignature := "e78e2223638cb60dbdbc88d23deb9b927ac41be7263ab38758605bac834dc25425705543707504bfef0802914cfa3f5f538fa308d1f9086211c420e7892ba2ba"

	data := "Postman-1578568851" + "200" + "10000.00" + "VT-server-HJMpl9HLr_ntOKt5mRONdmKj"
	signature := ComputeSHA512(data)

	t.Log(clSignature)
	t.Log(signature)

	assert.Equal(t, clSignature, signature)
}
