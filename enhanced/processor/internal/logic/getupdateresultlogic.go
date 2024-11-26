package logic

import (
	"context"
	"errors"

	"finishy1995/device-manager/enhanced/processor/internal/svc"
	"finishy1995/device-manager/enhanced/processor/pb/processor"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpdateResultLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpdateResultLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpdateResultLogic {
	return &GetUpdateResultLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Benchmark use
func (l *GetUpdateResultLogic) GetUpdateResult(in *processor.GetUpdateResultReq) (*processor.GetUpdateResultResp, error) {
	if task, ok := taskList[in.JobId]; ok {
		return &processor.GetUpdateResultResp{
			End:                        task.success,
			DeviceNumber:               task.deviceNumber,
			DeviceParamNumber:          task.deviceParamNumber,
			Thread:                     task.thread,
			Seconds:                    task.seconds,
			SuccessDeviceCount:         task.successDeviceCount,
			AverageLatencyMicroseconds: task.latencyMicroseconds / int64(task.successDeviceCount),
		}, nil
	}

	return &processor.GetUpdateResultResp{
		End: false,
	}, errors.New("task not found")
}
