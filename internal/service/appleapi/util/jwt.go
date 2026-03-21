package util

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"strings"
)

// JWTHeader JWT头部
type JWTHeader struct {
	Typ string   `json:"typ"`
	Alg string   `json:"alg"`
	Kid string   `json:"kid,omitempty"`
	X5c []string `json:"x5c,omitempty"`
}

// JWTAlgorithm JWT算法配置
var JWTAlgorithm = struct {
	Name          string
	Hash          crypto.Hash
	EllipticCurve elliptic.Curve
}{
	Name:          "ES256",
	Hash:          crypto.SHA256,
	EllipticCurve: elliptic.P256(),
}

// RequiredJWTHeaders 必需的JWT头部字段
var RequiredJWTHeaders = []string{"alg", "x5c"}

// CreateJWT 生成JWT
func CreateJWT(keyID string, privateKey *ecdsa.PrivateKey, payload interface{}) (string, error) {
	// 创建头部
	header := JWTHeader{
		Typ: "JWT",
		Alg: JWTAlgorithm.Name,
		Kid: keyID,
	}

	// 编码头部
	headerBytes, err := json.Marshal(header)
	if err != nil {
		return "", fmt.Errorf("failed to marshal header: %w", err)
	}
	headerEncoded := base64URLEncode(headerBytes)

	// 编码payload
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal payload: %w", err)
	}
	payloadEncoded := base64URLEncode(payloadBytes)

	// 生成签名
	signedData := []byte(fmt.Sprintf("%s.%s", headerEncoded, payloadEncoded))
	hash := sha256.Sum256(signedData)
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		return "", fmt.Errorf("failed to sign data: %w", err)
	}

	// 转换签名格式
	signatureBytes, err := encodeASN1Sequence(r, s)
	if err != nil {
		return "", fmt.Errorf("failed to encode signature: %w", err)
	}
	signatureEncoded := base64URLEncode(signatureBytes)

	// 组合JWT
	jwt := fmt.Sprintf("%s.%s.%s", headerEncoded, payloadEncoded, signatureEncoded)
	return jwt, nil
}

// verifyX509Chain 验证X509证书链
func verifyX509Chain(chain []string, rootCertificate []byte) error {
	if len(chain) < 3 {
		return errors.New("invalid certificate chain length")
	}

	// 解析证书链
	certificates := make([]*x509.Certificate, 0, len(chain))
	for _, certPEM := range chain {
		cert, err := parseCertificate([]byte(certPEM))
		if err != nil {
			return fmt.Errorf("failed to parse certificate: %w", err)
		}
		certificates = append(certificates, cert)
	}

	// 验证证书链
	cert := certificates[0]
	intermediate := certificates[1]
	appleRoot := certificates[2]

	// 验证证书是否由中间证书签名
	if err := cert.CheckSignatureFrom(intermediate); err != nil {
		return fmt.Errorf("certificate signature verification failed: %w", err)
	}

	// 验证中间证书是否由根证书签名
	if err := intermediate.CheckSignatureFrom(appleRoot); err != nil {
		return fmt.Errorf("intermediate certificate signature verification failed: %w", err)
	}

	// 如果提供了自定义根证书，验证苹果根证书是否由自定义根证书签名
	if rootCertificate != nil {
		customRoot, err := parseCertificate(rootCertificate)
		if err != nil {
			return fmt.Errorf("failed to parse custom root certificate: %w", err)
		}

		if err := appleRoot.CheckSignatureFrom(customRoot); err != nil {
			return fmt.Errorf("root certificate signature verification failed: %w", err)
		}
	}

	return nil
}

// verifySignature 验证JWT签名
func verifySignature(signedData, signatureBytes []byte, certPEM string) error {
	// 解析证书
	cert, err := parseCertificate([]byte(certPEM))
	if err != nil {
		return fmt.Errorf("failed to parse certificate: %w", err)
	}

	// 提取公钥
	publicKey, ok := cert.PublicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("certificate does not contain ECDSA public key")
	}

	// 解析签名
	var rs asn1Signature
	if _, err := asn1.Unmarshal(signatureBytes, &rs); err != nil {
		return fmt.Errorf("failed to unmarshal signature: %w", err)
	}

	// 计算哈希
	hash := sha256.Sum256(signedData)

	// 验证签名
	if !ecdsa.Verify(publicKey, hash[:], rs.R, rs.S) {
		return errors.New("signature verification failed")
	}

	return nil
}

// parseCertificate 解析证书
func parseCertificate(certData []byte) (*x509.Certificate, error) {
	// 尝试直接解析DER格式证书
	cert, err := x509.ParseCertificate(certData)
	if err == nil {
		return cert, nil
	}

	// 尝试解析PEM格式证书
	block, _ := pem.Decode(certData)
	if block != nil {
		cert, err = x509.ParseCertificate(block.Bytes)
		if err == nil {
			return cert, nil
		}
	}

	return nil, fmt.Errorf("failed to parse certificate: %w", err)
}

// asn1Signature ASN1签名结构
type asn1Signature struct {
	R, S *big.Int
}

// encodeASN1Sequence 编码R和S为ASN1序列
func encodeASN1Sequence(r, s *big.Int) ([]byte, error) {
	sig := asn1Signature{R: r, S: s}
	return asn1.Marshal(sig)
}

// base64URLEncode Base64URL编码
func base64URLEncode(data []byte) string {
	s := base64.URLEncoding.EncodeToString(data)
	// 移除填充
	s = strings.TrimRight(s, "=")
	return s
}

// base64URLDecode Base64URL解码
func base64URLDecode(s string) ([]byte, error) {
	// 添加填充
	for len(s)%4 != 0 {
		s += "="
	}
	return base64.URLEncoding.DecodeString(s)
}

// FormatPEM 格式化证书为PEM格式
func FormatPEM(cert string) string {
	if strings.HasPrefix(cert, "-----BEGIN CERTIFICATE-----") {
		return cert
	}

	// 添加PEM头部和尾部
	pemLines := make([]string, 0)
	pemLines = append(pemLines, "-----BEGIN CERTIFICATE-----")

	// 每64个字符换行
	for i := 0; i < len(cert); i += 64 {
		end := i + 64
		if end > len(cert) {
			end = len(cert)
		}
		pemLines = append(pemLines, cert[i:end])
	}

	pemLines = append(pemLines, "-----END CERTIFICATE-----")
	return strings.Join(pemLines, "\n")
}
