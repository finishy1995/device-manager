package logic

import (
	"context"
	"encoding/json"
	"time"

	"finishy1995/device-manager/enhanced/http/internal/svc"
	"finishy1995/device-manager/enhanced/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMetricsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMetricsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMetricsLogic {
	return &GetMetricsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMetricsLogic) GetMetrics(req *types.GetMetricsRequest) (resp *types.GetMetricsResponse, err error) {
	resp = &types.GetMetricsResponse{
		Code: 403,
	}
	if req.SN == "" {
		return
	}
	if req.EndTime == 0 {
		req.EndTime = time.Now().Unix()
	}

	// because customer has many different kinds of devices, and the data schema is different, so they need to create multiple tables to store the data.
	// For DEMO, we just use two tables to store the data - camera and sweeper
	if isCamera(req.SN) {
		data, err := l.svcCtx.DeviceCameraDataModel.FindByDeviceSnTimeRange(l.ctx, req.SN, req.StartTime, req.EndTime)
		if err != nil {
			l.Logger.Errorf("FindByDeviceSnTimeRange error: %v", err)
			return resp, nil
		}
		jsonData, err := json.Marshal(data)
		if err != nil {
			l.Logger.Errorf("json.Marshal error: %v", err)
			return resp, nil
		}
		resp.Code = 200
		resp.Data = string(jsonData)
	} else {
		data, err := l.svcCtx.DeviceSweeperDataModel.FindByDeviceSnTimeRange(l.ctx, req.SN, req.StartTime, req.EndTime)
		if err != nil {
			l.Logger.Errorf("FindByDeviceSnTimeRange error: %v", err)
			return resp, nil
		}
		jsonData, err := json.Marshal(data)
		if err != nil {
			l.Logger.Errorf("json.Marshal error: %v", err)
			return resp, nil
		}
		resp.Code = 200
		resp.Data = string(jsonData)
	}

	return
}

func isCamera(sn string) bool {
	// camera SN format -- `SN-1xxxxxxx`
	// sweeper SN format -- `SN-2xxxxxxx`
	// DEMO only
	return sn[3] == '1'
}
