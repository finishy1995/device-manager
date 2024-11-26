package logic

import (
	"context"
	"fmt"

	"finishy1995/device-manager/enhanced/http/internal/svc"
	"finishy1995/device-manager/enhanced/http/internal/types"
	"finishy1995/device-manager/enhanced/processor/processorclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateDemoMetadataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenerateDemoMetadataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateDemoMetadataLogic {
	return &GenerateDemoMetadataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateDemoMetadataLogic) GenerateDemoMetadata(req *types.GenerateDemoMetadataReq) (resp *types.GenerateDemoMetadataResp, err error) {
	fmt.Println(l.svcCtx.Processor)
	originResp, err := l.svcCtx.Processor.GenerateDemoMetadata(l.ctx, &processorclient.GenerateDemoMetadataReq{
		DeviceNumber:      int32(req.DeviceNumber),
		DeviceParamNumber: int32(req.DeviceParamNumber),
	})
	resp = &types.GenerateDemoMetadataResp{
		Code: int(originResp.Code),
	}

	return
}
