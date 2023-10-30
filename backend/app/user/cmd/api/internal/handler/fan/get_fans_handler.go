package fan

import (
	"net/http"

	"TokTik/app/user/cmd/api/internal/logic/fan"
	"TokTik/app/user/cmd/api/internal/svc"
	"TokTik/app/user/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetFansHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetFansgRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := fan.NewGetFansLogic(r.Context(), svcCtx)
		resp, err := l.GetFans(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
