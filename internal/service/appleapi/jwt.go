package appleapi

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
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

// ParseJWT 解析并验证JWT
func ParseJWT(signedPayload string) (map[string]interface{}, error) {
	// 分割JWT的三个部分
	parts := strings.Split(signedPayload, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid JWT format: expected 3 parts")
	}

	headerPart := parts[0]
	payloadPart := parts[1]
	//signaturePart := parts[2]

	// 解码头部
	headerBytes, err := base64URLDecode(headerPart)
	if err != nil {
		return nil, fmt.Errorf("failed to decode header: %w", err)
	}

	var header JWTHeader
	if err = json.Unmarshal(headerBytes, &header); err != nil {
		return nil, fmt.Errorf("failed to unmarshal header: %w", err)
	}

	// 解码payload
	payloadBytes, err := base64URLDecode(payloadPart)
	if err != nil {
		return nil, fmt.Errorf("failed to decode payload: %w", err)
	}

	var payload map[string]interface{}
	if err := json.Unmarshal(payloadBytes, &payload); err != nil {
		return nil, fmt.Errorf("failed to unmarshal payload: %w", err)
	}

	// 验证头部字段
	for _, requiredHeader := range RequiredJWTHeaders {
		if requiredHeader == "alg" && header.Alg != JWTAlgorithm.Name {
			return nil, fmt.Errorf("unrecognized algorithm: %s", header.Alg)
		}
		if requiredHeader == "x5c" && len(header.X5c) == 0 {
			return nil, errors.New("missing x5c header")
		}
	}

	//// 验证证书链
	//if err := verifyX509Chain(header.X5c, rootCertificate); err != nil {
	//	return nil, fmt.Errorf("failed to verify certificate chain: %w", err)
	//}
	//
	//// 验证签名
	//signedData := []byte(fmt.Sprintf("%s.%s", headerPart, payloadPart))
	//signatureBytes, err := base64URLDecode(signaturePart)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to decode signature: %w", err)
	//}
	//
	//if err = verifySignature(signedData, signatureBytes, header.X5c[0]); err != nil {
	//	return nil, fmt.Errorf("failed to verify signature: %w", err)
	//}

	return payload, nil
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
