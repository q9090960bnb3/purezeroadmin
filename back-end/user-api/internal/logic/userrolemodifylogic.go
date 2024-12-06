package logic

import (
	"context"
	"errors"

	"backend/user-api/helper"
	"backend/user-api/internal/svc"
	"backend/user-api/internal/types"
	"backend/utls/arrutil"
	"backend/utls/jsonutil"

	jsoniter "github.com/json-iterator/go"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserRoleModifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 增加角色
func NewUserRoleModifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRoleModifyLogic {
	return &UserRoleModifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRoleModifyLogic) UserRoleModify(req *types.UserRoleModifyReq) (resp string, err error) {
	_, mTbRole, err := helper.GetAuthsInfos(l.ctx, l.svcCtx)
	if err != nil {
		return "", err
	}

	tbRole, ok := mTbRole[req.Id]
	if !ok {
		return "", errors.New("无对应校色权限")
	}
	mActiveID := make(map[int64]struct{})
	for _, id := range req.Ids {
		mActiveID[id] = struct{}{}
	}

	rolePermissions, err := jsonutil.ToArray[string](tbRole.Permissions)
	if err != nil {
		return "", err
	}

	routers, err := l.svcCtx.TbRouterModel.FindAll(l.ctx)
	if err != nil {
		return "", err
	}

	for _, router := range routers {
		routerID := router.Id
		var getPass1, getPass2 bool
		var getPassPermission string
		getPass1, _ = l.svcCtx.Enforcer.Enforce(tbRole.Code, router.Path, "get")
		for _, permission := range rolePermissions {
			getPass, _ := l.svcCtx.Enforcer.Enforce(permission, router.Path, "get")
			if getPass {
				getPass2 = true
				getPassPermission = permission
				break
			}
		}
		if getPass1 || getPass2 {
			// 如果通过，但修改中不存在，则删除对应权限
			if _, ok := mActiveID[routerID]; !ok {
				_, err = l.svcCtx.Enforcer.DeletePermissionForUser(tbRole.Code, router.Path)
				if err != nil {
					return "", err
				}
				if getPass2 {
					rolePermissions = arrutil.RemoveItem(rolePermissions, getPassPermission)
					tbRole.Permissions, err = jsoniter.MarshalToString(rolePermissions)
					if err != nil {
						return "", err
					}
					err = l.svcCtx.TbRoleModel.Update(l.ctx, tbRole)
					if err != nil {
						return "", err
					}
				}
			}
		} else {
			if _, ok := mActiveID[routerID]; ok {
				// 如果通过，但修改中存在，则添加对应权限
				_, err = l.svcCtx.Enforcer.AddPermissionForUser(tbRole.Code, router.Path, "get")
				if err != nil {
					return "", err
				}
			}
		}
	}

	return "ok", nil
}
