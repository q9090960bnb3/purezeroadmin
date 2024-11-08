package logic

import (
	"context"

	"backend/user-api/helper"
	"backend/user-api/internal/svc"
	"backend/user-api/internal/types"
	"backend/utls/arrutil"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRouterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户登出
func NewUserRouterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRouterLogic {
	return &UserRouterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRouterLogic) UserRouter(req *types.UserRouterReq) (resp []*types.RouterData, err error) {
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

	isAdmin := arrutil.Contains(roles, "admin")

	return l.GetRecursionRoutersByParentID(0, isAdmin, roles, permissions)
}

func (l *UserRouterLogic) GetRouterByID(id int64, isAdmin bool, roles, permissions []string) (routerData *types.RouterData, err error) {
	router, err := l.svcCtx.TbRouterModel.FindOne(l.ctx, id)
	if err != nil {
		return nil, err
	}

	return helper.RouterToData(l.svcCtx, router, roles, permissions)
}

func (l *UserLoginLogic) GetRoutersByParentID(parentID int64, isAdmin bool, roles, permissions []string) (routerDatas []*types.RouterData, err error) {
	routers, err := l.svcCtx.TbRouterModel.FindAllFromParentID(l.ctx, parentID)
	if err != nil {
		return nil, err
	}

	for _, v := range routers {
		routerData, err := helper.RouterToData(l.svcCtx, v, roles, permissions)
		if err != nil {
			return nil, err
		}
		routerDatas = append(routerDatas, routerData)
	}

	return routerDatas, nil
}

func (l *UserRouterLogic) UpdateRouterData(routerData *types.RouterData, id int64, isAdmin bool, roles, permissions []string) (err error) {
	routers, err := l.svcCtx.TbRouterModel.FindAllFromParentID(l.ctx, id)
	if err != nil {
		return err
	}

	for _, v := range routers {
		child, err := l.GetRecursionRouterByID(v.Id, isAdmin, roles, permissions)
		if err != nil {
			return err
		}
		if child == nil {
			continue
		}
		routerData.Children = append(routerData.Children, child)
	}
	return nil
}

func (l *UserRouterLogic) GetRecursionRouterByID(id int64, isAdmin bool, roles, permissions []string) (routerData *types.RouterData, err error) {
	routerData, err = l.GetRouterByID(id, isAdmin, roles, permissions)
	if err != nil {
		return nil, err
	}

	err = l.UpdateRouterData(routerData, id, isAdmin, roles, permissions)
	return routerData, err
}

func (l *UserRouterLogic) GetRecursionRoutersByParentID(parentID int64, isAdmin bool, roles, permissions []string) (routerDatas []*types.RouterData, err error) {
	routers, err := l.svcCtx.TbRouterModel.FindAllFromParentID(l.ctx, parentID)
	if err != nil {
		return nil, err
	}

	for _, v := range routers {
		routerData, err := helper.RouterToData(l.svcCtx, v, roles, permissions)
		if err != nil {
			return nil, err
		}
		err = l.UpdateRouterData(routerData, v.Id, isAdmin, roles, permissions)
		if err != nil {
			return nil, err
		}
		routerDatas = append(routerDatas, routerData)
	}

	return routerDatas, nil
}
