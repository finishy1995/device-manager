package logic

import (
	"context"

	"finishy1995/device-manager/legacy/processor/internal/svc"
	"finishy1995/device-manager/legacy/processor/pb/processor"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMetadataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMetadataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMetadataLogic {
	return &UpdateMetadataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Benchmark use
func (l *UpdateMetadataLogic) UpdateMetadata(in *processor.UpdateMetadataReq) (*processor.UpdateMetadataResp, error) {
	t := NewTask(in.DeviceNumber, in.DeviceParamNumber, in.Thread, in.Seconds)
	go func() {
		t.Run(l.svcCtx.Config.DataSource)
	}()

	return &processor.UpdateMetadataResp{
		JobId: t.id,
	}, nil
}
