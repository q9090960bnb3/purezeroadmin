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
	"fmt"
)

func GetUserIDFromContext(ctx context.Context) (int64, error) {

	uid, ok := ctx.Value(global.CtxJwtUserIDKey).(json.Number)
	if !ok {
		return 0, fmt.Errorf("jwt has no userID")
	}

	return uid.Int64()
}

func GetAuths(ctx context.Context, svcCtx *svc.ServiceContext, tbUser *models.TbUser) (roles, permissions []string, err error) {
	roles, err = jsonutil.ToArray[string](tbUser.Roles)
	if err != nil {
		return nil, nil, err
	}

	for _, role := range roles {
		tbRole, err := svcCtx.TbRoleModel.FindOne(ctx, role)
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

func RouterToData(router *models.TbRouter, isAdmin bool, roles, permissions []string) (routerData *types.RouterData, err error) {
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

	if router.MetaRoles.Valid {
		routerData.Meta.Roles, err = jsonutil.ToArray[string](router.MetaRoles.String)
		if err != nil {
			return nil, err
		}
		if !isAdmin && !arrutil.ContainsAny(roles, routerData.Meta.Roles...) {
			return nil, nil
		}
	}

	if router.MetaAuths.Valid {
		routerData.Meta.Auths, err = jsonutil.ToArray[string](router.MetaAuths.String)
		if err != nil {
			return nil, err
		}
		if !isAdmin && !arrutil.ContainsAny(permissions, routerData.Meta.Auths...) {
			return nil, nil
		}
	}

	return routerData, nil
}
