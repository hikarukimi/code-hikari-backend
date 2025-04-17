package handler

import (
	"net/http"

	"code-hikari/user/api/internal/logic"
	"code-hikari/user/api/internal/svc"
	"code-hikari/user/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func MobileLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MobileLoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewMobileLoginLogic(r.Context(), svcCtx)
		resp, err := l.MobileLogin(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
