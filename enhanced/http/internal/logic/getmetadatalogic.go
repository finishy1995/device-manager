package logic

import (
	"context"
	"encoding/json"

	"finishy1995/device-manager/enhanced/http/internal/svc"
	"finishy1995/device-manager/enhanced/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMetadataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMetadataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMetadataLogic {
	return &GetMetadataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMetadataLogic) GetMetadata(req *types.GetMetadataRequest) (resp *types.GetMetadataResponse, err error) {
	resp = &types.GetMetadataResponse{
		Code: 403,
	}
	if req.SN == "" {
		return
	}
	metadata, err := l.svcCtx.DeviceMetadataModel.FindByDeviceSn(l.ctx, req.SN)
	if err != nil {
		l.Logger.Error("get metadata failed", err)
		return
	}
	if metadata == nil {
		resp.Code = 404
		return
	}

	bytes, err := json.Marshal(metadata)
	if err != nil {
		l.Logger.Error("marshal metadata failed", err)
		return
	}

	resp.Data = string(bytes)
	resp.Code = 200
	return
}
