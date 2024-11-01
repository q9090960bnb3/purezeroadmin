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
				Path:    "/login",
				Handler: userLoginHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 刷新token
				Method:  http.MethodPost,
				Path:    "/refresh-token",
				Handler: UserRefreshTokenHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 获取路由
				Method:  http.MethodGet,
				Path:    "/get-async-routes",
				Handler: userRouterHandler(serverCtx),
			},
		},
	)
}
