package handler

import (
	"net/http"

	"TokTik/app/video/api/cmd/internal/logic"
	"TokTik/app/video/api/cmd/internal/svc"
	"TokTik/app/video/api/cmd/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadVideoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploapReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUploadVideoLogic(r.Context(), svcCtx)
		resp, err := l.UploadVideo(&req, r)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
