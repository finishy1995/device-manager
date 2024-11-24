package handler

import (
	"net/http"

	"finishy1995/device-manager/legacy/http/internal/logic"
	"finishy1995/device-manager/legacy/http/internal/svc"
	"finishy1995/device-manager/legacy/http/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetRandomUpdateResultHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUpdateResultReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetRandomUpdateResultLogic(r.Context(), svcCtx)
		resp, err := l.GetRandomUpdateResult(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
