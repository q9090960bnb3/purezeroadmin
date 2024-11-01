package logic

import (
	"context"

	"backend/user-api/internal/svc"
	"backend/user-api/internal/types"

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

func (l *UserRouterLogic) UserRouter(req *types.UserRouterReq) (resp *types.UserRouterResp, err error) {
	return &types.UserRouterResp{
		Base: types.Base{
			Success: true,
		},
		Data: []types.RouterData{
			{
				Path: "/permission",
				Meta: types.Meta{
					Title: "权限管理",
					Icon:  "ep:lollipop",
					Rank:  10,
				},
				Children: []types.RouterData{
					{
						Path: "/permission/page/index",
						Name: "PermissionPage",
						Meta: types.Meta{
							Title: "页面权限",
							Roles: []string{"admin", "common"},
						},
					},
					{
						Path: "/permission/button",
						Meta: types.Meta{
							Title: "按钮权限",
							Roles: []string{"admin", "common"},
						},
						Children: []types.RouterData{
							{
								Path:      "/permission/button/router",
								Component: "permission/button/index",
								Name:      "PermissionButtonRouter",
								Meta: types.Meta{
									Title: "路由返回按钮权限",
									Auths: []string{
										"permission:btn:add",
										"permission:btn:edit",
										"permission:btn:delete",
									},
								},
							},
							{
								Path:      "/permission/button/login",
								Component: "permission/button/perms",
								Name:      "PermissionButtonLogin",
								Meta: types.Meta{
									Title: "登录返回按钮权限",
								},
							},
						},
					},
				},
			},
		},
	}, nil
}
