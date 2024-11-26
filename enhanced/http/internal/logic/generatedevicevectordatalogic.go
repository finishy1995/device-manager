package logic

import (
	"context"

	"finishy1995/device-manager/enhanced/http/internal/svc"
	"finishy1995/device-manager/enhanced/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateDeviceVectorDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenerateDeviceVectorDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateDeviceVectorDataLogic {
	return &GenerateDeviceVectorDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateDeviceVectorDataLogic) GenerateDeviceVectorData(req *types.GenerateDeviceVectorDataReq) (resp *types.GenerateDeviceVectorDataResp, err error) {
	resp = &types.GenerateDeviceVectorDataResp{Code: 200}
	// 1. 从数据库中，按批次把数据加载出来

	// 2. 生成向量数据

	// 3. 保存向量数据到数据库中

	return
}
