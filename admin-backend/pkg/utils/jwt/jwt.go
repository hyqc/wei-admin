package jwt

import (
	"crypto/rsa"
	"encoding/base64"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"strconv"
	"time"
)

type Auth struct {
	// publicKey base64 encoded, used by JWT
	publicKey *rsa.PublicKey
	// privateKey base64 encoded, used by JWT
	privateKey *rsa.PrivateKey
}

var (
	encoder = base64.StdEncoding
)

// NewAuth 实例化Auth
// publicKey string 经过base64编码的pem公钥字符串值
// privateKey string 经过base64编码的pem私钥字符串值
// return *Auth, error
func NewAuth(publicKey, privateKey string) (auth *Auth, err error) {
	auth = &Auth{}
	auth.privateKey, err = auth.parseBase64PrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	auth.publicKey, err = auth.parseBase64PublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	return auth, nil
}

func (a *Auth) parseBase64PrivateKey(p string) (*rsa.PrivateKey, error) {
	body, err := encoder.DecodeString(p)
	if err != nil {
		return nil, err
	}
	return jwt.ParseRSAPrivateKeyFromPEM(body)
}

func (a *Auth) parseBase64PublicKey(p string) (*rsa.PublicKey, error) {
	body, err := encoder.DecodeString(p)
	if err != nil {
		return nil, err
	}
	return jwt.ParseRSAPublicKeyFromPEM(body)
}

// Generate 生成Token
func (a *Auth) Generate(data ClaimsData) (*Token, error) {
	claims := newCustomClaims(data)
	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(a.privateKey)
	if err != nil {
		return nil, err
	}
	return &Token{
		Token:     token,
		CreatedAt: claims.createdAt,
		ExpiryAt:  claims.ExpiresAt.Local(),
	}, nil
}

// Inspect 验证Token
func (a *Auth) Inspect(t string) (*CustomClaims, error) {
	result := &CustomClaims{}
	token, err := jwt.ParseWithClaims(t, result, func(t *jwt.Token) (any, error) {
		return a.publicKey, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid == false {
		return nil, errors.New("token valid failed")
	}
	return result, err
}

type Token struct {
	// The actual token
	Token string `json:"token"`
	// Time of token creation
	CreatedAt time.Time `json:"created_at"`
	// Time of token expiry
	ExpiryAt time.Time `json:"expiry_at"`
}

type CustomClaims struct {
	jwt.RegisteredClaims
	ClaimsData
}

type ClaimsData struct {
	AccountId     int32     // 账号ID
	ExpireSeconds int64     //过期秒数
	Meta          any       //其他元数据
	uuid          uint64    // 唯一ID
	expiredAt     time.Time // 过期秒数
	createdAt     time.Time
}

func newCustomClaims(data ClaimsData) *CustomClaims {
	data.expiredAt = time.Now().Add(time.Duration(data.ExpireSeconds) * time.Second)
	data.createdAt = time.Now()
	data.uuid = uint64(uuid.New().ID())
	return &CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(data.expiredAt),
			IssuedAt:  jwt.NewNumericDate(data.createdAt),
			ID:        strconv.FormatUint(data.uuid, 10),
		},
		ClaimsData: data,
	}
}

func (c *CustomClaims) GetExpiredAt() time.Time {
	return c.expiredAt
}

func (c *CustomClaims) GetCreatedAt() time.Time {
	return c.createdAt
}

func (c *CustomClaims) GetUUID() uint64 {
	return c.uuid
}
