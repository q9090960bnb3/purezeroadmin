package handler

import (
	"net/http"

	"backend/user-api/internal/logic"
	"backend/user-api/internal/svc"
	"backend/user-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 用户登出
func userRouterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserRouterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserRouterLogic(r.Context(), svcCtx)
		resp, err := l.UserRouter(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
