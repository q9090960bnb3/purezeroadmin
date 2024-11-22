package logic

import (
	"context"
	"time"

	"backend/user-api/helper"
	"backend/user-api/internal/svc"
	"backend/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRolePatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改角色某个属性
func NewUserRolePatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRolePatchLogic {
	return &UserRolePatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRolePatchLogic) UserRolePatch(req *types.UserRolePatchReq) (resp string, err error) {
	_, err = helper.CheckAdmin(l.ctx, l.svcCtx)
	if err != nil {
		return "", err
	}
	src, err := l.svcCtx.TbRoleModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return "", err
	}
	isModify := false
	if req.Code != nil && src.Code != *req.Code {
		src.Code = *req.Code
		isModify = true
	}
	if req.Name != nil && src.Name != *req.Name {
		src.Name = *req.Name
		isModify = true
	}
	if req.Remark != nil && src.Remark != *req.Remark {
		src.Remark = *req.Remark
		isModify = true
	}
	if req.Status != nil && src.Status != *req.Status {
		src.Status = *req.Status
		isModify = true
	}
	if isModify {
		src.UpdateTs = time.Now().UnixMilli()
	}

	err = l.svcCtx.TbRoleModel.Update(l.ctx, src)
	if err != nil {
		return "", err
	}

	return "ok", nil
}
