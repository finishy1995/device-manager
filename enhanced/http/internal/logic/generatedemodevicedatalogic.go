package logic

import (
	"context"

	"finishy1995/device-manager/enhanced/http/internal/svc"
	"finishy1995/device-manager/enhanced/http/internal/types"
	"finishy1995/device-manager/enhanced/processor/processorclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateDemoDeviceDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenerateDemoDeviceDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateDemoDeviceDataLogic {
	return &GenerateDemoDeviceDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateDemoDeviceDataLogic) GenerateDemoDeviceData(req *types.GenerateDemoDeviceDataReq) (resp *types.GenerateDemoDeviceDataResp, err error) {
	originResp, err := l.svcCtx.Processor.GenerateDemoDeviceData(l.ctx, &processorclient.GenerateDemoDeviceDataReq{
		DeviceNumber: int32(req.DeviceNumber),
		StartTime:    req.StartTime,
		EndTime:      req.EndTime,
	})
	resp = &types.GenerateDemoDeviceDataResp{
		Code: int(originResp.Code),
	}

	return
}
