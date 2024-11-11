package logic

import (
	"context"

	"backend/user-api/helper"
	"backend/user-api/internal/svc"
	"backend/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取菜单
func NewUserMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserMenuLogic {
	return &UserMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserMenuLogic) UserMenu(req *types.UserRoleMenuReq) (resp []*types.UserRoleMenu, err error) {
	userID, err := helper.GetUserIDFromContext(l.ctx)
	if err != nil {
		return nil, err
	}

	tbUser, err := l.svcCtx.TbUserModel.FindOne(l.ctx, userID)
	if err != nil {
		return nil, err
	}

	roles, permissions, err := helper.GetAuths(l.ctx, l.svcCtx, tbUser)
	if err != nil {
		return nil, err
	}

	err = l.GetMenuByParentID(0, roles, permissions, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *UserMenuLogic) GetMenuByParentID(parentID int64, roles, permissions []string, roleMenus *[]*types.UserRoleMenu) (err error) {
	routers, err := l.svcCtx.TbRouterModel.FindAllFromParentID(l.ctx, parentID)
	if err != nil {
		return err
	}

	for _, v := range routers {
		if pass := helper.RouterPass(l.svcCtx, v, roles, permissions); pass {
			*roleMenus = append(*roleMenus, &types.UserRoleMenu{
				ParentId: v.ParentId,
				Id:       v.Id,
				MenuType: v.MenuType,
				Title:    v.MetaTitle,
			})
			err = l.GetMenuByParentID(v.Id, roles, permissions, roleMenus)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
