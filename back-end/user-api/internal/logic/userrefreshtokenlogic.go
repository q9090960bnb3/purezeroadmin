package logic

import (
	"context"

	"backend/user-api/internal/svc"
	"backend/user-api/internal/types"

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
	if req.RefreshToken != "" {
		return &types.UserRefreshTokenResp{
			Base: types.Base{
				Success: true,
			},
			Data: types.UserRefreshTokenData{
				AccessToken:  "",
				RefreshToken: "",
				Expires:      "",
			},
		}, nil
	}
	return &types.UserRefreshTokenResp{}, nil
}
