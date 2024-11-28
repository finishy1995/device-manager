package logic

import (
	"context"
	"errors"
	"strconv"

	"finishy1995/device-manager/enhanced/http/internal/svc"
	"finishy1995/device-manager/enhanced/http/internal/types"
	"finishy1995/device-manager/enhanced/model"

	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type UpdateMetadataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMetadataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMetadataLogic {
	return &UpdateMetadataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMetadataLogic) UpdateMetadata(req *types.UpdateMetadataRequest) (resp *types.UpdateMetadataResponse, err error) {
	resp = &types.UpdateMetadataResponse{
		Code: 403,
	}
	if req.SN == "" || req.Params == nil || len(req.Params) == 0 {
		return
	}
	// Create a new DeviceMetadata instance
	deviceMetadata := &model.DeviceMetadata{
		DeviceSn: req.SN,
		Params:   make(map[string]model.Param),
	}
	// Populate the Params map
	for k, v := range req.Params {
		_, errAtoi := strconv.Atoi(k)
		if errAtoi != nil {
			return resp, errors.New("param type must be int")
		}
		deviceMetadata.Params[k] = model.Param{
			PV: v,
			CT: time.Now(),
			UT: time.Now(),
		}
	}
	// Prepare data for upsert
	data := []*model.DeviceMetadata{deviceMetadata}
	// Perform the upsert operation
	_, err = l.svcCtx.DeviceMetadataModel.Upsert(l.ctx, data)
	if err == nil {
		resp.Code = 200
	}
	return
}
