package crypto

import (
	"crypto/rsa"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	message := []byte("{\"user_id\":1,\"exp\":1634025600}")

	privateKey, err := GenerateKeyPair()
	assert.NoError(t, err)

	privateKeyPEM := PemEncodePrivateKey(privateKey)
	fmt.Printf("Private key:\n%s\n", privateKeyPEM)

	publicKey := privateKey.Public().(*rsa.PublicKey)
	publicKeyPEM, err := PemEncodePublicKey(publicKey)
	assert.NoError(t, err)
	fmt.Printf("Public key:\n%s\n", publicKeyPEM)

	ciphertext, err := EncryptWithPublicKey(message, publicKey)
	assert.NoError(t, err)

	fmt.Printf("Ciphertext: %x\n", ciphertext)
	plaintext, err := DecryptWithPrivateKey(ciphertext, privateKey)
	assert.NoError(t, err)

	fmt.Printf("Plaintext: %s\n", plaintext)

	assert.Equal(t, message, plaintext)
}

func TestKey(t *testing.T) {
	privateKey, err := GenerateKeyPair()
	if err != nil {
		panic(err)
	}
	privateKey2, err := GenerateKeyPair()
	assert.NoError(t, err)

	privateKeyPEM := PemEncodePrivateKey(privateKey)
	fmt.Printf("Private key:\n%s\n", privateKeyPEM)

	privateKeyPEM2 := PemEncodePrivateKey(privateKey2)
	fmt.Printf("Private key2:\n%s\n", privateKeyPEM2)

	assert.NotEqual(t, privateKeyPEM, privateKeyPEM2)
}

func TestEncodeDecodeBase64(t *testing.T) {
	message := fmt.Sprintf("key_%d", time.Now().Unix())
	encoded := EncodeBASE64URL(message)
	fmt.Printf("Encoded: %s\n", encoded)
	decoded, err := DecodeBASE64(encoded)
	assert.NoError(t, err)
	fmt.Printf("Decoded: %s\n", decoded)
	assert.Equal(t, message, decoded)
}

func TestSHA512(t *testing.T) {
	clSignature := "e78e2223638cb60dbdbc88d23deb9b927ac41be7263ab38758605bac834dc25425705543707504bfef0802914cfa3f5f538fa308d1f9086211c420e7892ba2ba"

	data := "Postman-1578568851" + "200" + "10000.00" + "VT-server-HJMpl9HLr_ntOKt5mRONdmKj"
	signature := ComputeSHA512(data)

	fmt.Println(clSignature)
	fmt.Println(signature)

	assert.Equal(t, clSignature, signature)
}
