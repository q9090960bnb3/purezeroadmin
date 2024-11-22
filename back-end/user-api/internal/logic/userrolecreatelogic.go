package logic

import (
	"context"
	"time"

	"backend/user-api/helper"
	"backend/user-api/internal/svc"
	"backend/user-api/internal/types"
	"backend/user-api/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRoleCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 增加角色
func NewUserRoleCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRoleCreateLogic {
	return &UserRoleCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRoleCreateLogic) UserRoleCreate(req *types.UserRoleCreateReq) (resp string, err error) {
	_, err = helper.CheckAdmin(l.ctx, l.svcCtx)
	if err != nil {
		return "", err
	}
	_, err = l.svcCtx.TbRoleModel.Insert(l.ctx, &models.TbRole{
		Code:        req.Code,
		Name:        req.Name,
		Remark:      req.Remark,
		Permissions: "[]",
		CreateTs:    time.Now().UnixMilli(),
		UpdateTs:    time.Now().UnixMilli(),
	})
	if err != nil {
		return "", err
	}

	return "ok", nil
}
