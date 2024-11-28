package logic

import (
	"context"
	"fmt"

	"finishy1995/device-manager/enhanced/http/internal/svc"
	"finishy1995/device-manager/enhanced/http/internal/types"
	"finishy1995/device-manager/enhanced/model"

	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/mongo"
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

type DeviceVector struct {
	DeviceSn                string    `bson:"device_sn"`
	StartTime               int64     `bson:"start_time"`
	EndTime                 int64     `bson:"end_time"`
	VectorEuclideanDistance []float64 `bson:"vector_euclidean_distance"`
	VectorAngle             []float64 `bson:"vector_angle"`
}

func (l *GenerateDeviceVectorDataLogic) GenerateDeviceVectorData(req *types.GenerateDeviceVectorDataReq) (resp *types.GenerateDeviceVectorDataResp, err error) {
	resp = &types.GenerateDeviceVectorDataResp{Code: 200}

	go func(model model.DeviceCameraDataModel, startTime, endTime int64, client *mongo.Client) {
		index := 0
		ctx := context.Background()
		for {
			// 1. 从数据库中，按批次把数据加载出来
			sn := fmt.Sprintf("SN-%d%011d", 1, index)
			index++
			result, err := model.FindByDeviceSnTimeRange(ctx, sn, startTime, endTime)
			if err != nil {
				break
			}
			if len(result) == 0 {
				break
			}

			// 2. 生成向量数据
			vectorEuclideanDistance, vectorAngle := generateVector(result)
			logx.Infof("sn: %s, vectorEuclideanDistance: %v, vectorAngle: %v", sn, vectorEuclideanDistance, vectorAngle)

			// 3. 保存向量数据
			client.Database("test").Collection("device_vector_data").InsertOne(ctx, &DeviceVector{
				DeviceSn:                sn,
				StartTime:               startTime,
				EndTime:                 endTime,
				VectorEuclideanDistance: vectorEuclideanDistance,
				VectorAngle:             vectorAngle,
			})
		}

	}(l.svcCtx.DeviceCameraDataModel, req.StartTime, req.EndTime, l.svcCtx.MongoClient)

	return
}

func generateVector(result []*model.DeviceCameraData) ([]float64, []float64) {
	x := 0.0
	y := 0.0
	z := 0.0
	vectorEuclideanDistance := make([]float64, 0, len(result))
	vectorAngle := make([]float64, 0, len(result))
	for i := 0; i < len(result); i++ {
		if i != 0 {
			vectorEuclideanDistance = append(vectorEuclideanDistance, EuclideanDistanceWithNormalize(x, y, z, result[i].RotationX, result[i].RotationY, result[i].RotationZ))
			vectorAngle = append(vectorAngle, AngleBetweenVectorsWithNormalize(x, y, z, result[i].RotationX, result[i].RotationY, result[i].RotationZ))
		}
		x = result[i].RotationX
		y = result[i].RotationY
		z = result[i].RotationZ
	}
	return vectorEuclideanDistance, vectorAngle
}
