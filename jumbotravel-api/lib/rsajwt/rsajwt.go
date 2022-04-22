package rsajwt

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"io"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var (
	userTTL time.Duration = 7 * 24 * time.Hour         // 7 Days
	apiTTL  time.Duration = 100 * 365 * 24 * time.Hour // 100 years
)

type AllowPolicy struct {
	Resources []string          `json:"resources"`
	Extra     map[string]string `json:"extra"`
}

type Claims struct {
	AllowPolicy
	TokenType   string `json:"type"`
	SubjectType string `json:"subtype"`

	jwt.StandardClaims
}

// Verifier uses a rsa public key to verify a JWT token
type Verifier struct {
	key    *rsa.PublicKey
	rsaJWT *jwt.SigningMethodRSA
}

// Signer uses a rsa private key to build a JWT token
type Signer struct {
	rsaJWT *jwt.SigningMethodRSA
	key    *rsa.PrivateKey
}

type JWT struct {
	Token  string
	Claims Claims
}

func NewVerifierFromPublicKeyFile(publicKeyPath string) (*Verifier, error) {
	file, err := os.Open(publicKeyPath)
	if err != nil {
		return nil, err
	}

	return NewVerifierFromPublicKey(file)
}

func NewVerifierFromPublicKey(publicKey io.Reader) (*Verifier, error) {
	pubPEMData, err := io.ReadAll(publicKey)
	if err != nil {
		return nil, err
	}

	block, rest := pem.Decode(pubPEMData)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, errors.New("failed to decode PEM block containing public key")
	}

	if len(rest) != 0 {
		return nil, errors.New("passed public key contains more than just the public key")
	}

	pkixPublicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	key, isRsaPublicKey := pkixPublicKey.(*rsa.PublicKey)
	if !isRsaPublicKey {
		return nil, errors.New("passed public key is not of type rsa")
	}

	return &Verifier{
		key:    key,
		rsaJWT: jwt.SigningMethodRS256,
	}, nil
}

func ShorterVerifierFromPublicKeyFile(publicKeyPath string) (*Verifier, error) {
	file, err := os.Open(publicKeyPath)
	if err != nil {
		return nil, err
	}

	pubPEMData, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(pubPEMData)
	if err != nil {
		return nil, err
	}

	return &Verifier{
		key:    key,
		rsaJWT: jwt.SigningMethodRS256,
	}, nil
}

func NewSignerFromPrivateKeyFile(privateKeyPath string) (*Signer, error) {
	file, err := os.Open(privateKeyPath)
	if err != nil {
		return nil, err
	}

	privatePEMData, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	block, rest := pem.Decode(privatePEMData)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("failed to decode PEM block containing private key " + block.Type)
	}

	if len(rest) != 0 {
		return nil, errors.New("passed public key contains more than just the private key")
	}

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return &Signer{
		key:    key,
		rsaJWT: jwt.SigningMethodRS256,
	}, nil
}

func NewSignerFromBytes(privateKey []byte) (*Signer, error) {
	block, rest := pem.Decode(privateKey)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("failed to decode PEM block containing private key")
	}

	if len(rest) != 0 {
		return nil, errors.New("passed public key contains more than just the private key")
	}

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return &Signer{
		key:    key,
		rsaJWT: jwt.SigningMethodRS256,
	}, nil
}

func ShorterSignerFromPrivateKeyFile(privateKeyPath string) (*Signer, error) {
	file, err := os.Open(privateKeyPath)
	if err != nil {
		return nil, err
	}
	return ShorterSignerFromPrivateKey(file)
}

func ShorterSignerFromPrivateKey(privateKey io.Reader) (*Signer, error) {
	privatePEMData, err := io.ReadAll(privateKey)
	if err != nil {
		return nil, err
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(privatePEMData)
	if err != nil {
		return nil, err
	}

	return &Signer{
		key:    key,
		rsaJWT: jwt.SigningMethodRS256,
	}, nil
}

func (s *Signer) Sign(allowPolicy AllowPolicy, subject, subjetType, tokenType string) (JWT, error) {
	now := time.Now().UTC()

	var ttl time.Duration
	switch tokenType {
	case "user":
		ttl = userTTL
	case "api", "dynamic":
		ttl = apiTTL
	default:
		return JWT{}, errors.New("wrong tokenType, allowed values are user and api")
	}

	claims := Claims{
		AllowPolicy: allowPolicy,
		TokenType:   tokenType,
		SubjectType: subjetType,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now.Unix(),
			ExpiresAt: now.Add(ttl).Unix(),
			Subject:   subject,
			Id:        uuid.NewString(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	ss, err := token.SignedString(s.key)
	if err != nil {
		return JWT{}, err
	}

	jwt := JWT{
		Token:  ss,
		Claims: claims,
	}

	return jwt, nil
}

func (v *Verifier) GetVerifiedValue(token string, env string) (Claims, error) {
	unixNow := time.Now().UTC().Unix()

	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return Claims{}, errors.New("Invalid API Key")
	}

	err := v.rsaJWT.Verify(strings.Join(parts[0:2], "."), parts[2], v.key)
	if err != nil {
		return Claims{}, err
	}
	claimsBytes, err := jwt.DecodeSegment(parts[1])
	if err != nil {
		return Claims{}, err
	}

	dec := json.NewDecoder(bytes.NewBuffer(claimsBytes))

	c := Claims{}
	err = dec.Decode(&c)

	if unixNow > c.ExpiresAt {
		return Claims{}, errors.New("Expired Token")
	}
	if unixNow < c.IssuedAt {
		return Claims{}, errors.New("Expired Token")
	}

	return c, err
}
