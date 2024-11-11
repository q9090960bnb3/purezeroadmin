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

func (l *UserRoleLogic) UserRole(req *types.UserRoleReq) (resp *types.UserRoleResp, err error) {
	list, total, err := l.svcCtx.TbRoleModel.FindList(l.ctx, req.Name, req.Code, req.Status, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	resp = &types.UserRoleResp{
		List:  make([]*types.UserRoleData, len(list)),
		Total: total,
	}

	for i, tbRole := range list {
		resp.List[i] = &types.UserRoleData{
			Id:         tbRole.Id,
			Code:       tbRole.Code,
			Name:       tbRole.Name,
			Status:     tbRole.Status,
			Remark:     tbRole.Remark,
			CreateTime: tbRole.CreateTs,
			UpdateTime: tbRole.UpdateTs,
		}
	}

	return resp, nil
}
