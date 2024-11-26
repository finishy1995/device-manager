package svc

import (
	"finishy1995/device-manager/enhanced/http/internal/config"
	"finishy1995/device-manager/enhanced/model"
	"finishy1995/device-manager/enhanced/processor/processorclient"

	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config                 config.Config
	DeviceMetadataModel    model.DeviceMetadataModel
	DeviceCameraDataModel  model.DeviceCameraDataModel
	DeviceSweeperDataModel model.DeviceSweeperDataModel
	Processor              processorclient.Processor
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	rpcConn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{c.Etcd},
			Key:   c.Processor,
		},
	})
	return &ServiceContext{
		Config:                 c,
		DeviceMetadataModel:    model.NewDeviceMetadataModel(conn),
		DeviceCameraDataModel:  model.NewDeviceCameraDataModel(conn),
		DeviceSweeperDataModel: model.NewDeviceSweeperDataModel(conn),
		Processor:              processorclient.NewProcessor(rpcConn),
	}
}
