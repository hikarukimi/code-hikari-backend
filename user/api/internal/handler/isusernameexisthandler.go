package handler

import (
	"net/http"

	"code-hikari/user/api/internal/logic"
	"code-hikari/user/api/internal/svc"
	"code-hikari/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func IsUsernameExistHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IsUsernameExistRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewIsUsernameExistLogic(r.Context(), svcCtx)
		resp, err := l.IsUsernameExist(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
