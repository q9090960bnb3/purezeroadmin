package handler

import (
	"net/http"

	"backend/hello-api/internal/logic"
	"backend/hello-api/internal/svc"
	"backend/hello-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 你好，世界
func helloInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HelloReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewHelloInfoLogic(r.Context(), svcCtx)
		resp, err := l.HelloInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
