// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package handler

import (
	"net/http"

	"backend/user-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// 用户登录
				Method:  http.MethodPost,
				Path:    "/api/login",
				Handler: userLoginHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 刷新token
				Method:  http.MethodPost,
				Path:    "/api/refresh-token",
				Handler: UserRefreshTokenHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 获取路由
				Method:  http.MethodGet,
				Path:    "/api/get-async-routes",
				Handler: userRouterHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 获取路由
				Method:  http.MethodGet,
				Path:    "/api/role",
				Handler: userRoleHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 获取菜单
				Method:  http.MethodGet,
				Path:    "/api/role-menu",
				Handler: userMenuHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 获取菜单详情
				Method:  http.MethodPost,
				Path:    "/api/role-menu-ids",
				Handler: userMenuIDHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
