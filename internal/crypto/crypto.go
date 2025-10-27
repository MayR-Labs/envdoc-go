package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"

	"golang.org/x/crypto/pbkdf2"
)

const (
	saltSize   = 32
	iterations = 10000
	keySize    = 32
)

// Encrypt encrypts data using AES-256-CBC with PBKDF2
func Encrypt(data []byte, password string) (string, error) {
	// Generate salt
	salt := make([]byte, saltSize)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("failed to generate salt: %w", err)
	}

	// Derive key using PBKDF2
	key := pbkdf2.Key([]byte(password), salt, iterations, keySize, sha256.New)

	// Create cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %w", err)
	}

	// Pad data
	paddedData := pkcs7Pad(data, aes.BlockSize)

	// Create IV
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return "", fmt.Errorf("failed to generate IV: %w", err)
	}

	// Encrypt
	ciphertext := make([]byte, len(paddedData))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, paddedData)

	// Combine salt + iv + ciphertext
	result := append(salt, iv...)
	result = append(result, ciphertext...)

	// Encode to base64
	return base64.StdEncoding.EncodeToString(result), nil
}

// Decrypt decrypts data using AES-256-CBC with PBKDF2
func Decrypt(encryptedData, password string) ([]byte, error) {
	// Decode from base64
	data, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode data: %w", err)
	}

	if len(data) < saltSize+aes.BlockSize {
		return nil, fmt.Errorf("encrypted data is too short")
	}

	// Extract salt, IV, and ciphertext
	salt := data[:saltSize]
	iv := data[saltSize : saltSize+aes.BlockSize]
	ciphertext := data[saltSize+aes.BlockSize:]

	// Derive key using PBKDF2
	key := pbkdf2.Key([]byte(password), salt, iterations, keySize, sha256.New)

	// Create cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher: %w", err)
	}

	// Decrypt
	plaintext := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, ciphertext)

	// Unpad data
	unpaddedData, err := pkcs7Unpad(plaintext, aes.BlockSize)
	if err != nil {
		return nil, fmt.Errorf("decryption failed (wrong password?): %w", err)
	}

	return unpaddedData, nil
}

// HashSHA256 returns the SHA256 hash of data
func HashSHA256(data []byte) string {
	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash)
}

// pkcs7Pad pads data to the specified block size
func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padText := make([]byte, padding)
	for i := range padText {
		padText[i] = byte(padding)
	}
	return append(data, padText...)
}

// pkcs7Unpad removes padding from data
func pkcs7Unpad(data []byte, blockSize int) ([]byte, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("data is empty")
	}

	padding := int(data[len(data)-1])
	if padding > blockSize || padding > len(data) {
		return nil, fmt.Errorf("invalid padding")
	}

	for i := len(data) - padding; i < len(data); i++ {
		if data[i] != byte(padding) {
			return nil, fmt.Errorf("invalid padding")
		}
	}

	return data[:len(data)-padding], nil
}

// EncodeBase64 encodes data to base64
func EncodeBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// DecodeBase64 decodes data from base64
func DecodeBase64(encoded string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(encoded)
}

// GenerateRandomBytes generates random bytes of the specified length
func GenerateRandomBytes(length int) ([]byte, error) {
	bytes := make([]byte, length)
	if _, err := io.ReadFull(rand.Reader, bytes); err != nil {
		return nil, err
	}
	return bytes, nil
}
