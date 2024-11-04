package logic

import (
	"context"
	"errors"

	"backend/user-api/internal/svc"
	"backend/user-api/internal/types"

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
	if req.UserName == "admin" && req.Password == "admin123" {
		return &types.UserLoginResp{
			Base: types.Base{
				Success: true,
			},
			Data: types.UserLoginData{
				Avatar:       "https://avatars.githubusercontent.com/u/44761321",
				Username:     "admin",
				Nickname:     "小铭",
				Roles:        []string{"admin"},
				Permissions:  []string{"*:*:*"},
				AccessToken:  "eyJhbGciOiJIUzUxMiJ9.admin",
				RefreshToken: "eyJhbGciOiJIUzUxMiJ9.adminRefresh",
				Expires:      "2030/10/30 00:00:00",
			},
		}, nil
	}

	if req.UserName == "common" && req.Password == "123456" {
		return &types.UserLoginResp{
			Base: types.Base{
				Success: true,
			},
			Data: types.UserLoginData{
				Avatar:       "https://avatars.githubusercontent.com/u/52823142",
				Username:     "common",
				Nickname:     "小林",
				Roles:        []string{"common"},
				Permissions:  []string{"permission:btn:add", "permission:btn:edit"},
				AccessToken:  "eyJhbGciOiJIUzUxMiJ9.common",
				RefreshToken: "eyJhbGciOiJIUzUxMiJ9.commonRefresh",
				Expires:      "2030/10/30 00:00:00",
			},
		}, nil
	}

	return nil, errors.New("登陆失败")
}
