package helper

import (
	"backend/user-api/global"
	"backend/user-api/internal/svc"
	"backend/user-api/internal/types"
	"backend/user-api/models"
	"backend/utls/arrutil"
	"backend/utls/jsonutil"
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

func GetUserIDFromContext(ctx context.Context) (int64, error) {
	uid, ok := ctx.Value(global.CtxJwtUserIDKey).(json.Number)
	if !ok {
		return 0, fmt.Errorf("jwt has no userID")
	}

	return uid.Int64()
}

func GetUser(ctx context.Context, svcCtx *svc.ServiceContext) (isAdmin bool, tbUser *models.TbUser, err error) {
	userID, err := GetUserIDFromContext(ctx)
	if err != nil {
		return false, nil, err
	}

	tbUser, err = svcCtx.TbUserModel.FindOne(ctx, userID)
	if err != nil {
		return false, nil, err
	}
	return tbUser.Username == "admin", tbUser, nil
}

func CheckAdmin(ctx context.Context, svcCtx *svc.ServiceContext) (tbUser *models.TbUser, err error) {
	isAdmin, _, err := GetUser(ctx, svcCtx)
	if err != nil {
		return nil, err
	}
	if !isAdmin {
		return nil, errors.New("必须Admin权限")
	}
	return tbUser, nil
}

func GetAuthsInfos(ctx context.Context, svcCtx *svc.ServiceContext) (tbUser *models.TbUser, mTbRole map[int64]*models.TbRole, err error) {
	isAdmin, tbUser, err := GetUser(ctx, svcCtx)
	if err != nil {
		return nil, nil, err
	}
	mTbRole = make(map[int64]*models.TbRole)

	if isAdmin {
		tbRoles, err := svcCtx.TbRoleModel.FindAll(ctx)
		if err != nil {
			return nil, nil, err
		}

		for _, tbRole := range tbRoles {
			mTbRole[tbRole.Id] = tbRole
		}
		return tbUser, mTbRole, nil
	}

	roles, err := jsonutil.ToArray[string](tbUser.Roles)
	if err != nil {
		return nil, nil, err
	}

	for _, role := range roles {
		tbRole, err := svcCtx.TbRoleModel.FindOneByCode(ctx, role)
		if err != nil {
			return nil, nil, err
		}

		mTbRole[tbRole.Id] = tbRole
	}

	return tbUser, mTbRole, nil
}

func GetAuths(ctx context.Context, svcCtx *svc.ServiceContext, tbUser *models.TbUser) (roles, permissions []string, err error) {
	roles, err = jsonutil.ToArray[string](tbUser.Roles)
	if err != nil {
		return nil, nil, err
	}

	for _, role := range roles {
		tbRole, err := svcCtx.TbRoleModel.FindOneByCode(ctx, role)
		if err != nil {
			return nil, nil, err
		}

		rolePermissions, err := jsonutil.ToArray[string](tbRole.Permissions)
		if err != nil {
			return nil, nil, err
		}

		permissions = arrutil.UniqueConcat(permissions, rolePermissions)
	}

	return roles, permissions, nil
}

func RouterPass(svcCtx *svc.ServiceContext, router *models.TbRouter, roles, permissions []string) bool {
	for _, role := range roles {
		getPass, _ := svcCtx.Enforcer.Enforce(role, router.Path, "get")
		if getPass {
			return true
		}
	}
	for _, permission := range permissions {
		getPass, _ := svcCtx.Enforcer.Enforce(permission, router.Path, "get")
		if getPass {
			return true
		}
	}
	return false
}

func RouterToData(svcCtx *svc.ServiceContext, router *models.TbRouter, roles, permissions []string) (routerData *types.RouterData, err error) {
	routerData = &types.RouterData{
		Path:      router.Path,
		Name:      router.Name,
		Component: router.Component,
		Meta: types.Meta{
			Title: router.MetaTitle,
			Icon:  router.MetaIcon,
			Rank:  router.MetaRank,
		},
	}

	pass := RouterPass(svcCtx, router, roles, permissions)
	if !pass {
		return nil, nil
	}

	if router.MetaRoles.Valid {
		routerData.Meta.Roles, err = jsonutil.ToArray[string](router.MetaRoles.String)
		if err != nil {
			return nil, err
		}
	}

	if router.MetaAuths.Valid {
		routerData.Meta.Auths, err = jsonutil.ToArray[string](router.MetaAuths.String)
		if err != nil {
			return nil, err
		}
	}

	return routerData, nil
}
