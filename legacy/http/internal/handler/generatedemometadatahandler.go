package handler

import (
	"net/http"

	"finishy1995/device-manager/legacy/http/internal/logic"
	"finishy1995/device-manager/legacy/http/internal/svc"
	"finishy1995/device-manager/legacy/http/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GenerateDemoMetadataHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GenerateDemoMetadataReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGenerateDemoMetadataLogic(r.Context(), svcCtx)
		resp, err := l.GenerateDemoMetadata(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
