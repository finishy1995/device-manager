// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"finishy1995/device-manager/enhanced/http/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/generate/devicedata",
				Handler: GenerateDemoDeviceDataHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/generate/metadata",
				Handler: GenerateDemoMetadataHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/generate/vector",
				Handler: GenerateDeviceVectorDataHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/metadata",
				Handler: GetMetadataHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/metadata/update",
				Handler: UpdateMetadataHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/metrics",
				Handler: GetMetricsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/search/vector",
				Handler: SearchBadDeviceHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/update/metadata",
				Handler: RandomUpdateMetadataHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/update/result",
				Handler: GetRandomUpdateResultHandler(serverCtx),
			},
		},
	)
}
