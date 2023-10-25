package handler

import (
	"TokTik/app/video/api/cmd/internal/logic"
	"TokTik/app/video/api/cmd/internal/svc"
	"TokTik/app/video/api/cmd/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadVedioHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploapReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUploadVedioLogic(r.Context(), svcCtx)
		resp, err := l.UploadVedio(&req, r)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
