package logic

import (
	"context"

	"finishy1995/device-manager/enhanced/processor/internal/svc"
	"finishy1995/device-manager/enhanced/processor/pb/processor"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateDemoMetadataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateDemoMetadataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateDemoMetadataLogic {
	return &GenerateDemoMetadataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DEMO use
func (l *GenerateDemoMetadataLogic) GenerateDemoMetadata(in *processor.GenerateDemoMetadataReq) (*processor.GenerateDemoMetadataResp, error) {
	result := generateDemoMetadata(in.DeviceNumber, in.DeviceParamNumber)
	resp := &processor.GenerateDemoMetadataResp{
		Code: 400,
	}
	if len(result) > 0 {
		go func() {
			batchResult, err := l.svcCtx.DeviceMetadataModel.Upsert(context.Background(), result)
			if err != nil {
				l.Errorf("Upsert error: %v", err)
			} else {
				l.Infof("Upsert success: %v", batchResult.SuccessCount)
			}
		}()

		resp.Code = 200
	}

	return resp, nil
}
