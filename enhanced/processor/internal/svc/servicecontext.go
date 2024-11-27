package svc

import (
	"context"
	"finishy1995/device-manager/enhanced/model"
	"finishy1995/device-manager/enhanced/processor/internal/config"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ServiceContext struct {
	Config                 config.Config
	DeviceMetadataModel    model.DeviceMetadataModel
	DeviceCameraDataModel  model.DeviceCameraDataModel
	DeviceSweeperDataModel model.DeviceSweeperDataModel
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

	return &ServiceContext{
		Config:                 c,
		DeviceMetadataModel:    model.NewDeviceMetadataModel(client),
		DeviceCameraDataModel:  model.NewDeviceCameraDataModel(client),
		DeviceSweeperDataModel: model.NewDeviceSweeperDataModel(client),
	}
}

