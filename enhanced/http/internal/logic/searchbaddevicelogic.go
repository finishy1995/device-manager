package logic

import (
	"context"

	"finishy1995/device-manager/enhanced/http/internal/svc"
	"finishy1995/device-manager/enhanced/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchBadDeviceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchBadDeviceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchBadDeviceLogic {
	return &SearchBadDeviceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchBadDeviceLogic) SearchBadDevice(req *types.SearchBadDeviceReq) (resp *types.SearchBadDeviceResp, err error) {
	// todo: add your logic here and delete this line

	return
}
