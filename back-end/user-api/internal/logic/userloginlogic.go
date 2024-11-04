package logic

import (
	"context"
	"errors"
	"time"

	"backend/user-api/internal/svc"
	"backend/user-api/internal/types"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户登录
func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginReq) (resp *types.UserLoginResp, err error) {
	if (req.UserName == "admin" && req.Password == "admin123") || (req.UserName == "common" && req.Password == "123456") {

		tNow := time.Now()
		tExpire := tNow.Add(time.Second * time.Duration(l.svcCtx.Config.Auth.AccessExpire))
		accessToken, err := l.getJwtToken(req.UserName, l.svcCtx.Config.Auth.AccessSecret, tNow, tExpire)
		if err != nil {
			return nil, err
		}
		// 可免登陆1天
		tRefreshExpire := tExpire.AddDate(0, 0, 1)
		refreshToken, err := l.getJwtToken(req.UserName, l.svcCtx.Config.Auth.AccessSecret, tNow, tRefreshExpire)
		if err != nil {
			return nil, err
		}

		return &types.UserLoginResp{
			Base: types.Base{
				Success: true,
			},
			Data: types.UserLoginData{
				Avatar:       "https://avatars.githubusercontent.com/u/44761321",
				Username:     req.UserName,
				Nickname:     req.UserName + "_xxx",
				Roles:        []string{"admin"},
				Permissions:  []string{"*:*:*"},
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
				Expires:      tExpire.Format("2006/01/02 15:04:05"), //  "2030/10/30 00:00:00",
			},
		}, nil
	}

	return nil, errors.New("用户名或密码错误")
}

func (l *UserLoginLogic) getJwtToken(user, secretKey string, tNow, tExpire time.Time) (string, error) {

	claims := &jwt.RegisteredClaims{
		Issuer:    user,
		ExpiresAt: jwt.NewNumericDate(tExpire),
		IssuedAt:  jwt.NewNumericDate(tNow),
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
