package sec

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
	"path/filepath"
)

const Application = "fsm"

func SaveJWT(jwt string) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	plaintext := []byte(jwt)

	publicKey := &privateKey.PublicKey
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plaintext)
	if err != nil {
		panic(err)
	}

	dir, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	jwtPath := filepath.Join(dir, Application, "jwt")
	privateKeyPath := filepath.Join(dir, Application, "private.key.pem")

	secJwt, err := os.Create(jwtPath)
	defer secJwt.Close()
	if _, err = secJwt.Write(ciphertext); err != nil {
		return err
	}

	return savePrivateKey(privateKey, privateKeyPath)
}

func ReadJWT() ([]byte, error) {

	dir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}

	jwtPath := filepath.Join(dir, Application, "jwt")
	privateKeyPath := filepath.Join(dir, Application, "private.key.pem")

	ciphertext, err := os.ReadFile(jwtPath)
	if err != nil {
		return nil, err
	}

	key, err := loadPrivateKey(privateKeyPath)
	if err != nil {
		return nil, err
	}

	decryptedText, err := rsa.DecryptPKCS1v15(rand.Reader, key, ciphertext)
	if err != nil {
		return nil, err
	}

	return decryptedText, err
}

// 将私钥保存到磁盘
func savePrivateKey(privateKey *rsa.PrivateKey, filename string) error {
	keyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	pemBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: keyBytes,
	}
	return os.WriteFile(filename, pem.EncodeToMemory(pemBlock), 0600)
}

// 加载私钥
func loadPrivateKey(filename string) (*rsa.PrivateKey, error) {
	keyData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyData)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the private key")
	}

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return key, nil
}
