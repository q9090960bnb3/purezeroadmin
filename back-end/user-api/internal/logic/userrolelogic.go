package logic

import (
	"context"

	"backend/user-api/internal/svc"
	"backend/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取路由
func NewUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRoleLogic {
	return &UserRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRoleLogic) UserRole(req *types.UserRoleReq) (resp []*types.UserRoleResp, err error) {
	tbRoles, err := l.svcCtx.TbRoleModel.FindAll(l.ctx)
	if err != nil {
		return nil, err
	}

	for _, tbRole := range tbRoles {
		resp = append(resp, &types.UserRoleResp{
			Id:         tbRole.Id,
			Code:       tbRole.Code,
			Name:       tbRole.Name,
			Status:     int(tbRole.Status),
			Remark:     tbRole.Remark,
			CreateTime: tbRole.CreateTs,
			UpdateTime: tbRole.UpdateTs,
		})
	}

	return resp, nil
}
