// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"TokTik/app/video/cmd/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/upload_vedio",
				Handler: UploadVedioHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/api/v1/vedio"),
	)
}
