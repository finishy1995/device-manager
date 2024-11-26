package svc

import (
	"finishy1995/device-manager/enhanced/model"
	"finishy1995/device-manager/enhanced/processor/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config                 config.Config
	DeviceMetadataModel    model.DeviceMetadataModel
	DeviceCameraDataModel  model.DeviceCameraDataModel
	DeviceSweeperDataModel model.DeviceSweeperDataModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:                 c,
		DeviceMetadataModel:    model.NewDeviceMetadataModel(conn),
		DeviceCameraDataModel:  model.NewDeviceCameraDataModel(conn),
		DeviceSweeperDataModel: model.NewDeviceSweeperDataModel(conn),
	}
}
