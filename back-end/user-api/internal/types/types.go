// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type Base struct {
	Success bool `json:"success"`
}

type Meta struct {
	Title string   `json:"title"`
	Icon  string   `json:"icon,omitempty"`
	Rank  int64    `json:"rank,omitempty"`
	Roles []string `json:"roles,omitempty"`
	Auths []string `json:"auths,omitempty"`
}

type RouterData struct {
	Path      string       `json:"path"`
	Name      string       `json:"name,omitempty"`
	Component string       `json:"component,omitemty"`
	Meta      Meta         `json:"meta"`
	Children  []RouterData `json:"children,omitempty"`
}

type UserLoginData struct {
	Avatar       string   `json:"avatar"`
	Username     string   `json:"username"`
	Nickname     string   `json:"nickname"`
	Roles        []string `json:"roles"`
	Permissions  []string `json:"permissions"`
	AccessToken  string   `json:"accessToken"`
	RefreshToken string   `json:"refreshToken"`
	Expires      string   `json:"expires"`
}

type UserLoginReq struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResp struct {
	Base
	Data UserLoginData `json:"data"`
}

type UserRefreshTokenData struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	Expires      string `json:"expires"`
}

type UserRefreshTokenReq struct {
	RefreshToken string `json:"refreshToken"`
}

type UserRefreshTokenResp struct {
	Base
	Data UserRefreshTokenData `json:"data"`
}

type UserRouterReq struct {
}

type UserRouterResp struct {
	Base
	Data []RouterData `json:"data"`
}