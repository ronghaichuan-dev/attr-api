package appleapi

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
)

// Key 表示App Store服务器API的私钥
type Key struct {
	KeyID         string
	PrivateKey    *rsa.PrivateKey
	privateKeyPEM string
}

// NewKey 创建一个新的Key实例
func NewKey(keyID, privateKeyPEM string) (*Key, error) {
	// 解析PEM格式的私钥
	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil {
		return nil, errors.New("failed to decode PEM block containing private key")
	}

	// 解析RSA私钥
	privateKey, err := parseRSAPrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RSA private key: %w", err)
	}

	return &Key{
		KeyID:         keyID,
		PrivateKey:    privateKey,
		privateKeyPEM: privateKeyPEM,
	}, nil
}

// GetKeyID 获取Key ID
func (k *Key) GetKeyID() string {
	return k.KeyID
}

// GetPrivateKey 获取RSA私钥
func (k *Key) GetPrivateKey() *rsa.PrivateKey {
	return k.PrivateKey
}

// GetPrivateKeyPEM 获取PEM格式的私钥
func (k *Key) GetPrivateKeyPEM() string {
	return k.privateKeyPEM
}

// parseRSAPrivateKey 解析RSA私钥，支持PKCS#1和PKCS#8格式
func parseRSAPrivateKey(data []byte) (*rsa.PrivateKey, error) {
	// 尝试解析PKCS#1格式
	key, err := x509.ParsePKCS1PrivateKey(data)
	if err == nil {
		return key, nil
	}

	// 尝试解析PKCS#8格式
	key8, err := x509.ParsePKCS8PrivateKey(data)
	if err == nil {
		if rsaKey, ok := key8.(*rsa.PrivateKey); ok {
			return rsaKey, nil
		}
		return nil, errors.New("private key is not RSA")
	}

	return nil, fmt.Errorf("failed to parse private key: %w", err)
}
