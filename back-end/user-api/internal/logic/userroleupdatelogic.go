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

type UserRoleUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改角色
func NewUserRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRoleUpdateLogic {
	return &UserRoleUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRoleUpdateLogic) UserRoleUpdate(req *types.UserRoleUpdateReq) (resp string, err error) {
	_, err = helper.CheckAdmin(l.ctx, l.svcCtx)
	if err != nil {
		return "", err
	}
	sourceRoleModel, err := l.svcCtx.TbRoleModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return "", err
	}

	roleModel := &models.TbRole{
		Id:          req.Id,
		Code:        req.Code,
		Name:        req.Name,
		Status:      sourceRoleModel.Status,
		Remark:      req.Remark,
		Permissions: sourceRoleModel.Permissions,
		CreateTs:    sourceRoleModel.CreateTs,
		UpdateTs:    time.Now().UnixMilli(),
	}

	err = l.svcCtx.TbRoleModel.Update(l.ctx, roleModel)
	if err != nil {
		return "", err
	}

	return "ok", nil
}
