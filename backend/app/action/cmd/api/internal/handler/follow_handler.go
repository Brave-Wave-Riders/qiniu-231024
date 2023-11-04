package handler

import (
	"net/http"

	"TokTik/app/action/cmd/api/internal/logic"
	"TokTik/app/action/cmd/api/internal/svc"
	"TokTik/app/action/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func followHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FollowRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewFollowLogic(r.Context(), svcCtx)
		resp, err := l.Follow(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
