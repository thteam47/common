package jwtrepo

import "github.com/thteam47/common/entity"

type JwtRepository interface {
	Generate(tokenInfo *entity.TokenInfo) (string, error)
	Verify(accessToken string) (*entity.Claims, error)
	KeyJwt() string
}
