package handler

import (
	"net/http"

	"backend/user-api/internal/logic"
	"backend/user-api/internal/svc"
	"backend/user-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

	xhttp "github.com/zeromicro/x/http"
)

// 增加角色
func UserRoleModifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserRoleModifyReq
		if err := httpx.Parse(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserRoleModifyLogic(r.Context(), svcCtx)
		resp, err := l.UserRoleModify(&req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
