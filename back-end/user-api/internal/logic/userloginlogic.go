package logic

import (
	"context"
	"errors"
	"fmt"
	"time"

	"backend/user-api/global"
	"backend/user-api/internal/svc"
	"backend/user-api/internal/types"
	"backend/utls/jwtutil"

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
	var userID int64
	resp = &types.UserLoginResp{
		Username: req.UserName,
	}
	fmt.Println("req:", req)
	if req.UserName == "admin" && req.Password == "admin123" {
		userID = 1
		resp.Avatar = "https://avatars.githubusercontent.com/u/44761321"
		resp.Nickname = "小铭"
		resp.Roles = []string{"admin"}
		resp.Permissions = []string{"*:*:*"}
	} else if req.UserName == "common" && req.Password == "common123" {
		userID = 2
		resp.Avatar = "https://avatars.githubusercontent.com/u/52823142"
		resp.Nickname = "小林"
		resp.Roles = []string{"common"}
		resp.Permissions = []string{"permission:btn:add", "permission:btn:edit"}
	} else {
		return nil, errors.New("登陆失败")
	}

	tNow := time.Now()
	tExpire := tNow.Add(time.Second * time.Duration(l.svcCtx.Config.Auth.AccessExpire))

	mPayload := map[string]any{
		global.CtxJwtUserIDKey: userID,
	}

	accessToken, err := jwtutil.GetToken(l.svcCtx.Config.Auth.AccessSecret, tNow.Unix(), tExpire.Unix(), mPayload)
	if err != nil {
		return nil, err
	}

	refreshToken, err := jwtutil.GetToken(l.svcCtx.Config.Auth.AccessSecret, tNow.Unix(), tExpire.Unix()+86400, mPayload)
	if err != nil {
		return nil, err
	}

	resp.AccessToken = accessToken
	resp.RefreshToken = refreshToken
	resp.Expires = tExpire.Format("2006/01/02 15:04:05")

	return resp, nil
}
