package jwtImpl

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/thteam47/common-libs/errutil"
	"github.com/thteam47/common/entity"
	"github.com/thteam47/common/handler/jwt"
)

type JwtRepositoryImpl struct {
	jwtKey string
}

func NewJwtRepo(jwtKey string) jwtrepo.JwtRepository {
	return &JwtRepositoryImpl{
		jwtKey: jwtKey,
	}
}
func (inst *JwtRepositoryImpl) Generate(tokenInfo *entity.TokenInfo) (string, error) {
	claims := entity.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: int64(tokenInfo.Exp),
		},
		TokenInfo: tokenInfo,
	}

	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := claim.SignedString([]byte(inst.jwtKey))
	if err != nil {
		return "", errutil.Wrapf(err, "claim.SignedString")
	}
	return token, nil
}
func (inst *JwtRepositoryImpl) Verify(accessToken string) (*entity.Claims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&entity.Claims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return []byte(inst.jwtKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*entity.Claims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}

func (inst *JwtRepositoryImpl) KeyJwt() string {
	return inst.jwtKey
}
