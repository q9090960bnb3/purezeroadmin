package logic

import (
	"context"

	"backend/user-api/helper"
	"backend/user-api/internal/svc"
	"backend/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRoleDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除角色
func NewUserRoleDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRoleDeleteLogic {
	return &UserRoleDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRoleDeleteLogic) UserRoleDelete(req *types.UserRoleDeleteReq) (resp string, err error) {
	_, err = helper.CheckAdmin(l.ctx, l.svcCtx)
	if err != nil {
		return "", err
	}
	err = l.svcCtx.TbRoleModel.Delete(l.ctx, req.Id)
	if err != nil {
		return "", err
	}
	return "ok", nil
}
