package logic

import (
	"context"
	"errors"

	"backend/user-api/helper"
	"backend/user-api/internal/svc"
	"backend/user-api/internal/types"
	"backend/utls/jsonutil"

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
	_, mTbRole, err := helper.GetAuthsInfos(l.ctx, l.svcCtx)
	if err != nil {
		return nil, err
	}

	tbRole, ok := mTbRole[req.Id]
	if !ok {
		return nil, errors.New("无对应校色权限")
	}

	rolePermissions, err := jsonutil.ToArray[string](tbRole.Permissions)
	if err != nil {
		return nil, err
	}

	routers, err := l.svcCtx.TbRouterModel.FindAll(l.ctx)
	if err != nil {
		return nil, err
	}

	for _, router := range routers {
		if !helper.RouterPass(l.svcCtx, router, []string{tbRole.Code}, rolePermissions) {
			continue
		}
		resp = append(resp, router.Id)
	}

	return resp, nil
}
