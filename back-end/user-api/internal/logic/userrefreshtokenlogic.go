package logic

import (
	"context"
	"time"

	"backend/user-api/global"
	"backend/user-api/helper"
	"backend/user-api/internal/svc"
	"backend/user-api/internal/types"
	"backend/utls/jwtutil"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRefreshTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户登录
func NewUserRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRefreshTokenLogic {
	return &UserRefreshTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRefreshTokenLogic) UserRefreshToken(req *types.UserRefreshTokenReq) (resp *types.UserRefreshTokenResp, err error) {
	userID, err := helper.GetUserIDFromContext(l.ctx)
	if err != nil {
		return nil, err
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

	return &types.UserRefreshTokenResp{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Expires:      tExpire.Format("2006/01/02 15:04:05"),
	}, nil
}
