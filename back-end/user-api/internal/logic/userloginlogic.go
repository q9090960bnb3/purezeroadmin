package logic

import (
	"context"
	"errors"
	"time"

	"backend/user-api/global"
	"backend/user-api/helper"
	"backend/user-api/internal/svc"
	"backend/user-api/internal/types"
	"backend/utls/codeutil"
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

	tbUser, err := l.svcCtx.TbUserModel.FindOneByUsername(l.ctx, req.UserName)
	if err != nil {
		return nil, err
	}

	if tbUser.Password != codeutil.Md5Str(req.Password) {
		return nil, errors.New("用户或密码错误")
	}

	roles, permissions, err := helper.GetAuths(l.ctx, l.svcCtx, tbUser)
	if err != nil {
		return nil, err
	}

	tNow := time.Now()
	tExpire := tNow.Add(time.Second * time.Duration(l.svcCtx.Config.Auth.AccessExpire))

	mPayload := map[string]any{
		global.CtxJwtUserIDKey: tbUser.UserId,
	}

	accessToken, err := jwtutil.GetToken(l.svcCtx.Config.Auth.AccessSecret, tNow.Unix(), tExpire.Unix(), mPayload)
	if err != nil {
		return nil, err
	}

	refreshToken, err := jwtutil.GetToken(l.svcCtx.Config.Auth.AccessSecret, tNow.Unix(), tExpire.Unix()+86400, mPayload)
	if err != nil {
		return nil, err
	}

	return &types.UserLoginResp{
		Avatar:       tbUser.Avatar,
		Username:     tbUser.Username,
		Nickname:     tbUser.Nickname,
		Roles:        roles,
		Permissions:  permissions,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Expires:      tExpire.Format("2006/01/02 15:04:05"),
	}, nil
}
