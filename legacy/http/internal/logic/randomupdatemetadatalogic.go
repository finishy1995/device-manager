package logic

import (
	"context"

	"finishy1995/device-manager/legacy/http/internal/svc"
	"finishy1995/device-manager/legacy/http/internal/types"
	"finishy1995/device-manager/legacy/processor/pb/processor"

	"github.com/zeromicro/go-zero/core/logx"
)

type RandomUpdateMetadataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRandomUpdateMetadataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RandomUpdateMetadataLogic {
	return &RandomUpdateMetadataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RandomUpdateMetadataLogic) RandomUpdateMetadata(req *types.UpdateMetadataReq) (resp *types.UpdateMetadataResp, err error) {
	originResp, err := l.svcCtx.Processor.UpdateMetadata(l.ctx, &processor.UpdateMetadataReq{
		DeviceNumber:      int32(req.DeviceNumber),
		DeviceParamNumber: int32(req.DeviceParamNumber),
		Thread:            int32(req.Thread),
		Seconds:           int32(req.Seconds),
	})

	resp = &types.UpdateMetadataResp{
		JobId: int(originResp.JobId),
	}

	return
}
