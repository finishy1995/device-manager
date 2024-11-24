package logic

import (
	"context"

	"finishy1995/device-manager/legacy/http/internal/svc"
	"finishy1995/device-manager/legacy/http/internal/types"
	"finishy1995/device-manager/legacy/processor/pb/processor"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRandomUpdateResultLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRandomUpdateResultLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRandomUpdateResultLogic {
	return &GetRandomUpdateResultLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRandomUpdateResultLogic) GetRandomUpdateResult(req *types.GetUpdateResultReq) (resp *types.GetUpdateResultResp, err error) {
	originResp, err := l.svcCtx.Processor.GetUpdateResult(l.ctx, &processor.GetUpdateResultReq{
		JobId: int32(req.JobId),
	})
	resp = &types.GetUpdateResultResp{
		End:                        originResp.End,
		DeviceNumber:               int(originResp.DeviceNumber),
		DeviceParamNumber:          int(originResp.DeviceParamNumber),
		Thread:                     int(originResp.Thread),
		Seconds:                    int(originResp.Seconds),
		SuccessDeviceCount:         int(originResp.SuccessDeviceCount),
		AverageLatencyMicroseconds: originResp.AverageLatencyMicroseconds,
	}

	return
}
