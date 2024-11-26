package logic

import (
	"context"
	"fmt"
	"math/rand"

	"finishy1995/device-manager/legacy/model"
	"finishy1995/device-manager/legacy/processor/internal/svc"
	"finishy1995/device-manager/legacy/processor/pb/processor"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateDemoDeviceDataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateDemoDeviceDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateDemoDeviceDataLogic {
	return &GenerateDemoDeviceDataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DEMO use
func (l *GenerateDemoDeviceDataLogic) GenerateDemoDeviceData(in *processor.GenerateDemoDeviceDataReq) (*processor.GenerateDemoDeviceDataResp, error) {
	cameras := generateDemoCameraData(in.DeviceNumber/2, in.StartTime, in.EndTime)
	go func() {
		result, err := l.svcCtx.DeviceCameraDataModel.Upsert(context.Background(), cameras)
		if err != nil {
			l.Errorf("upsert camera data failed: %v", err)
		} else {
			l.Infof("upsert camera data success: %v", result)
		}
	}()
	sweepers := generateDemoSweeperData(in.DeviceNumber/2, in.StartTime, in.EndTime)
	go func() {
		result, err := l.svcCtx.DeviceSweeperDataModel.Upsert(context.Background(), sweepers)
		if err != nil {
			l.Errorf("upsert sweeper data failed: %v", err)
		} else {
			l.Infof("upsert sweeper data success: %v", result)
		}
	}()

	return &processor.GenerateDemoDeviceDataResp{}, nil
}

// For DEMO, assume device update frequency is 1 second
func generateDemoCameraData(deviceNumber int32, startTime int64, endTime int64) []*model.DeviceCameraData {
	result := make([]*model.DeviceCameraData, 0, deviceNumber)
	for i := 0; i < int(deviceNumber); i++ {
		batteryDec := rand.Intn(2) == 1
		batteryNow := rand.Intn(10000)
		rotationType := rand.Intn(100) // 模式确定：0-89 正常，90-94 渐进漂移，95-98 固定，99 非物理性旋转
		rotationX := getRandomRotation()
		rotationY := getRandomRotation()
		rotationZ := getRandomRotation()

		// 渐进漂移的角度增量
		driftX := 0.0
		driftY := 0.0
		driftZ := 0.0
		if rotationType >= 90 && rotationType <= 94 {
			driftX = rand.Float64()*0.5 - 0.25 // 每秒随机变化 -0.25 到 +0.25 度
			driftY = rand.Float64()*0.5 - 0.25
			driftZ = rand.Float64()*0.5 - 0.25
		}

		for j := startTime; j < endTime; j++ {
			data := &model.DeviceCameraData{}
			data.DeviceSn = fmt.Sprintf("SN-%d%011d", 1, i)
			data.Timestamp = j
			data.IsFixed = 0
			if batteryDec {
				batteryNow -= rand.Intn(10)
				if batteryNow < 0 {
					batteryNow = 0
				}
			} else {
				batteryNow += rand.Intn(10)
				if batteryNow > 10000 {
					batteryNow = 10000
				}
			}
			data.BatteryLevel = uint64(batteryNow)
			// 旋转角度生成逻辑
			switch {
			case rotationType >= 90 && rotationType <= 94: // 渐进漂移
				rotationX += driftX
				rotationY += driftY
				rotationZ += driftZ

				// 保证角度范围在 [0, 360)
				if rotationX < 0 {
					rotationX += 360
				}
				if rotationX >= 360 {
					rotationX -= 360
				}
				if rotationY < 0 {
					rotationY += 360
				}
				if rotationY >= 360 {
					rotationY -= 360
				}
				if rotationZ < 0 {
					rotationZ += 360
				}
				if rotationZ >= 360 {
					rotationZ -= 360
				}

				data.RotationX = rotationX
				data.RotationY = rotationY
				data.RotationZ = rotationZ

			case rotationType >= 95 && rotationType <= 98: // 固定不动
				data.RotationX = rotationX
				data.RotationY = rotationY
				data.RotationZ = rotationZ

			case rotationType == 99: // 非物理性旋转
				data.RotationX = getRandomRotation()
				data.RotationY = getRandomRotation()
				data.RotationZ = getRandomRotation()

			default: // 正常模式
				rotationX = getRandomChange(rotationX)
				rotationY = getRandomChange(rotationY)
				rotationZ = getRandomChange(rotationZ)

				data.RotationX = rotationX
				data.RotationY = rotationY
				data.RotationZ = rotationZ
			}

			// 是否固定标志
			if rotationType >= 95 && rotationType <= 98 {
				data.IsFixed = 1
			} else {
				data.IsFixed = 0
			}

			// 将生成的数据添加到结果集中
			result = append(result, data)
		}
	}
	return result
}

func generateDemoSweeperData(deviceNumber int32, startTime int64, endTime int64) []*model.DeviceSweeperData {
	result := make([]*model.DeviceSweeperData, 0, deviceNumber)
	for i := 0; i < int(deviceNumber); i++ {
		batteryDec := rand.Intn(2) == 1
		batteryNow := rand.Intn(10000)
		rotationType := rand.Intn(100) // 0-97: normal, 98: fixed, 99: broken
		rotationX := getRandomRotation()
		rotationY := getRandomRotation()
		for j := startTime; j < endTime; j++ {
			data := &model.DeviceSweeperData{}
			data.DeviceSn = fmt.Sprintf("SN-%d%011d", 1, i)
			data.Timestamp = j
			data.IsCharging = boolToInt64(!batteryDec)
			if batteryDec {
				batteryNow -= rand.Intn(10)
				if batteryNow < 0 {
					batteryNow = 0
				}
			} else {
				batteryNow += rand.Intn(10)
				if batteryNow > 10000 {
					batteryNow = 10000
				}
			}
			data.BatteryLevel = uint64(batteryNow)
			if rotationType > 98 {
				data.PositionX = getRandomRotation()
				data.PositionY = getRandomRotation()
			} else if rotationType == 98 {
				data.PositionX = rotationX
				data.PositionY = rotationY
			} else {
				rotationX = getRandomChange(rotationX)
				rotationY = getRandomChange(rotationY)
				data.PositionX = rotationX
				data.PositionY = rotationY
			}

			result = append(result, data)
		}
	}
	return result
}

func boolToInt64(b bool) int64 {
	if b {
		return 1
	}
	return 0
}

func getRandomRotation() float64 {
	return float64(rand.Intn(36000)) / 100.0
}

func getRandomChange(origin float64) float64 {
	new := origin + float64(rand.Intn(100)-50)/100.0
	if new < 0 {
		new += 360
	} else if new >= 360 {
		new -= 360
	}
	return new
}
