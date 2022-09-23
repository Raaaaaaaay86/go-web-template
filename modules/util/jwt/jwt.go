package jwt

import (
	"go-web-template/modules/constant/exception"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/wire"
)


//go:generate mockery --dir . --filename mock_crypt.go --name IJwtManager --output ../../mocks
type IJwtManager interface {
	Create() (token string, err error)
	Verify(token string) error
}

type JwtManager struct {
	signingKey string
}

var JwtManagerWireSet = wire.NewSet(
	wire.Bind(new(IJwtManager), new(JwtManager)),
	JwtManagerProvider,
)

func JwtManagerProvider() JwtManager {
	jwtManager := JwtManager{}
	jwtManager.setSigningKey("ddgdshkafgayaikshyvaksdghvdkasvgasdkyfgaysdjfggdshkafgayaikshyvaksdghvdkasvgasdkyfgaysdjfg")
	return jwtManager
}

func (jm JwtManager) Create() (tokenString string, err error) {
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(1 * time.Minute).Unix(),
		Issuer:    "OnlineMall",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(jm.signingKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (jm JwtManager) Verify(token string) error {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, exception.ErrJWTUnexpectedAlgorithm
		}

		return []byte(jm.signingKey), nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (jm *JwtManager) setSigningKey(key string) {
	jm.signingKey = key
}
