package crypto

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io"

	"golang.org/x/crypto/bcrypt"
)

// EncodeSHA1HMACBase64 : encrypt to SHA1HMAC input key, data String. Output to String in Base64 format
func EncodeSHA1HMACBase64(key string, data ...string) string {
	return EncodeBASE64(ComputeSHA1HMAC(key, data...))
}

// EncodeSHA1HMAC : encrypt to SHA1HMAC input key, data String. Output to String in Base16/Hex format
func EncodeSHA1HMAC(key string, data ...string) string {
	return fmt.Sprintf("%x", ComputeSHA1HMAC(key, data...))
}

// ComputeSHA1HMAC : encrypt to SHA1HMAC input key, data String. Output to String
func ComputeSHA1HMAC(key string, data ...string) string {
	h := hmac.New(sha1.New, []byte(key))
	for _, v := range data {
		io.WriteString(h, v)
	}
	return string(h.Sum(nil))
}

func EncodeSHA256HMACBase64(key string, data ...string) string {
	return EncodeBASE64(ComputeSHA256HMAC(key, data...))
}

func EncodeSHA256HMAC(key string, data ...string) string {
	return fmt.Sprintf("%x", ComputeSHA256HMAC(key, data...))
}

func ComputeSHA256HMAC(key string, data ...string) string {
	h := hmac.New(sha256.New, []byte(key))
	for _, v := range data {
		io.WriteString(h, v)
	}
	return string(h.Sum(nil))
}

func EncodeSHA512HMACBase64(key string, data ...string) string {
	return EncodeBASE64(ComputeSHA512HMAC(key, data...))
}

func EncodeSHA512HMAC(key string, data ...string) string {
	return fmt.Sprintf("%x", ComputeSHA512HMAC(key, data...))
}

func ComputeSHA512HMAC(key string, data ...string) string {
	h := hmac.New(sha512.New, []byte(key))
	for _, v := range data {
		io.WriteString(h, v)
	}
	return string(h.Sum(nil))
}

// EncodeMD5 : encrypt to MD5 input string, output to string
func EncodeMD5(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}

func EncodeMD5Base64(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	// return EncodeBASE64(hex.EncodeToString(h.Sum(nil)))
	return base64.StdEncoding.EncodeToString((h.Sum(nil)))
}

// EncodeBASE64 : Encrypt to Base64. Input string, output string
func EncodeBASE64(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}

// DecodeBASE64 : Decrypt Base64. Input string, output string
func DecodeBASE64(text string) (string, error) {
	byt, err := base64.StdEncoding.DecodeString(text)
	return string(byt), err
}

// EncodeBASE64URL : Encrypt to Base64URL. Input string, output text
func EncodeBASE64URL(text string) string {
	return base64.URLEncoding.EncodeToString([]byte(text))
}

// EncodeDES : Encrypt to DES. input string, output chiper
func EncodeDES(text string) (cipher.Block, error) {
	desKey, _ := hex.DecodeString(text)
	cipher, err := des.NewTripleDESCipher(desKey)
	return cipher, err
}

// EncodeSHA256: Encrypt to SHA256. input string, output text
func EncodeSHA256(text string) string {
	h := sha256.Sum256([]byte(text))
	return fmt.Sprintf("%x", h)
}

// EncodeSHA512 Encrypt to SHA512. input string, output text
func EncodeSHA512(text string) string {
	h := sha512.Sum512([]byte(text))
	return fmt.Sprintf("%x", h)
}

// HashPassword : Encrypt password using bcrypt. input string, output string
func HashPassword(password string, cost ...int) (string, error) {
	c := bcrypt.DefaultCost
	if len(cost) > 0 {
		c = cost[0]
	}

	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), c)
	return string(bcryptPassword), err
}

// CheckPasswordHash : Check password hash. input password, hash. output boolean
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateKeyPair : Generate RSA key pair. output private key, error
func GenerateKeyPair() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}

// PemEncodePrivateKey : Encode private key to PEM. input private key, output byte
func PemEncodePrivateKey(privateKey *rsa.PrivateKey) []byte {
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	return pem.EncodeToMemory(block)
}

// PemEncodePublicKey : Encode public key to PEM. input public key, output byte
func PemEncodePublicKey(publicKey *rsa.PublicKey) ([]byte, error) {
	bytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, err
	}

	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: bytes,
	}

	return pem.EncodeToMemory(block), nil
}

// EncryptWithPublicKey : Encrypt with public key. input message, public key. output byte, error
func EncryptWithPublicKey(message []byte, publicKey *rsa.PublicKey) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, publicKey, message)
}

// DecryptWithPrivateKey : Decrypt with private key. input ciphertext, private key. output byte, error
func DecryptWithPrivateKey(ciphertext []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
}

func ComputeSHA512(data string) string {
	h := sha512.New()
	h.Write([]byte(data))
	hashBytes := h.Sum(nil)
	return hex.EncodeToString(hashBytes)
}
