package svc

import (
	"finishy1995/device-manager/legacy/http/internal/config"
	"finishy1995/device-manager/legacy/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config              config.Config
	DeviceMetadataModel model.DeviceMetadataModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:              c,
		DeviceMetadataModel: model.NewDeviceMetadataModel(sqlx.NewMysql(c.DataSource)),
	}
}
