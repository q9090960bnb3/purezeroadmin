package logic

import (
	"context"

	"backend/user-api/internal/svc"
	"backend/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserMenuIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取菜单详情
func NewUserMenuIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserMenuIDLogic {
	return &UserMenuIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserMenuIDLogic) UserMenuID(req *types.UserRoleMenuIDReq) (resp []int64, err error) {
	// todo: add your logic here and delete this line

	return
}
