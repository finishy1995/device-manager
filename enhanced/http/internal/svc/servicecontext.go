package svc

import (
	"finishy1995/device-manager/enhanced/http/internal/config"
	"finishy1995/device-manager/enhanced/model"
	"finishy1995/device-manager/enhanced/processor/processorclient"

	"context"
	"fmt"
	"log"

	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ServiceContext struct {
	Config                 config.Config
	DeviceMetadataModel    model.DeviceMetadataModel
	DeviceCameraDataModel  model.DeviceCameraDataModel
	DeviceSweeperDataModel model.DeviceSweeperDataModel
	Processor              processorclient.Processor
	MongoClient            *mongo.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	// Set client options
	clientOptions := options.Client().ApplyURI(c.DataSource)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to connect to MongoDB: %v", err))
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to ping MongoDB: %v", err))
	}
	rpcConn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{c.Etcd},
			Key:   c.Processor,
		},
	})

	return &ServiceContext{
		Config:                 c,
		DeviceMetadataModel:    model.NewDeviceMetadataModel(client),
		DeviceCameraDataModel:  model.NewDeviceCameraDataModel(client),
		DeviceSweeperDataModel: model.NewDeviceSweeperDataModel(client),
		Processor:              processorclient.NewProcessor(rpcConn),
		MongoClient:            client,
	}
}
