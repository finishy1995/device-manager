// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: processor.proto

package server

import (
	"context"

	"finishy1995/device-manager/enhanced/processor/internal/logic"
	"finishy1995/device-manager/enhanced/processor/internal/svc"
	"finishy1995/device-manager/enhanced/processor/pb/processor"
)

type ProcessorServer struct {
	svcCtx *svc.ServiceContext
	processor.UnimplementedProcessorServer
}

func NewProcessorServer(svcCtx *svc.ServiceContext) *ProcessorServer {
	return &ProcessorServer{
		svcCtx: svcCtx,
	}
}

// DEMO use
func (s *ProcessorServer) GenerateDemoMetadata(ctx context.Context, in *processor.GenerateDemoMetadataReq) (*processor.GenerateDemoMetadataResp, error) {
	l := logic.NewGenerateDemoMetadataLogic(ctx, s.svcCtx)
	return l.GenerateDemoMetadata(in)
}

// Benchmark use
func (s *ProcessorServer) UpdateMetadata(ctx context.Context, in *processor.UpdateMetadataReq) (*processor.UpdateMetadataResp, error) {
	l := logic.NewUpdateMetadataLogic(ctx, s.svcCtx)
	return l.UpdateMetadata(in)
}

// Benchmark use
func (s *ProcessorServer) GetUpdateResult(ctx context.Context, in *processor.GetUpdateResultReq) (*processor.GetUpdateResultResp, error) {
	l := logic.NewGetUpdateResultLogic(ctx, s.svcCtx)
	return l.GetUpdateResult(in)
}

// DEMO use
func (s *ProcessorServer) GenerateDemoDeviceData(ctx context.Context, in *processor.GenerateDemoDeviceDataReq) (*processor.GenerateDemoDeviceDataResp, error) {
	l := logic.NewGenerateDemoDeviceDataLogic(ctx, s.svcCtx)
	return l.GenerateDemoDeviceData(in)
}
