package grpcauth

import "github.com/thteam47/common/entity"

type UserContext interface {
	GetUserId() string
	HasRoleAdmin() bool
	ServiceOrHasRoleAdmin() bool
	HasRole() string
	IsAuthenDone() bool
	SetAccessToken(token string)
	GetAccessToken() string
	GetTokenInfo() *entity.TokenInfo
}

type UserContextImpl struct {
	TokenInfo   *entity.TokenInfo
	accessToken string
}

func NewUserContext(tokenInfo *entity.TokenInfo) UserContext {
	return &UserContextImpl{
		TokenInfo: tokenInfo,
	}
}

func (inst *UserContextImpl) GetUserId() string {
	return inst.TokenInfo.UserId
}

func (inst *UserContextImpl) GetTokenInfo() *entity.TokenInfo {
	return inst.TokenInfo
}

func (inst *UserContextImpl) HasRoleAdmin() bool {
	return inst.TokenInfo.Role == "admin"
}

func (inst *UserContextImpl) ServiceOrHasRoleAdmin() bool {
	return inst.TokenInfo.Subject == "service" || inst.TokenInfo.Subject == "admin"
}
func (inst *UserContextImpl) HasRole() string {
	return inst.TokenInfo.Role
}

func (inst *UserContextImpl) IsAuthenDone() bool {
	return inst.TokenInfo.AuthenticationDone
}

func (inst *UserContextImpl) SetAccessToken(token string) {
	inst.accessToken = token
}

func (inst *UserContextImpl) GetAccessToken() string {
	return inst.accessToken
}
