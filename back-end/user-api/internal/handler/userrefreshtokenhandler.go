package handler

import (
	"net/http"

	"backend/user-api/internal/logic"
	"backend/user-api/internal/svc"
	"backend/user-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 用户登录
func UserRefreshTokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserRefreshTokenReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserRefreshTokenLogic(r.Context(), svcCtx)
		resp, err := l.UserRefreshToken(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
