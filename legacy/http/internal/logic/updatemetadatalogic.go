package logic

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"finishy1995/device-manager/legacy/http/internal/svc"
	"finishy1995/device-manager/legacy/http/internal/types"
	"finishy1995/device-manager/legacy/model"

	"github.com/zeromicro/go-zero/core/logx"
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
	data := make([]*model.DeviceMetadata, 0, len(req.Params))
	for k, v := range req.Params {
		t, errAtoi := strconv.Atoi(k)
		if errAtoi != nil {
			return resp, errors.New("param type must be int")
		}

		data = append(data, &model.DeviceMetadata{
			DeviceSn:   req.SN,
			ParamType:  int64(t),
			ParamValue: sql.NullString{String: v, Valid: true},
		})
	}
	// TODO: batch result error handler
	_, err = l.svcCtx.DeviceMetadataModel.Upsert(l.ctx, data)
	if err == nil {
		resp.Code = 200
	}

	return
}
