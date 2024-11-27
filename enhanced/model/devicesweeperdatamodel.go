package model

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var _ DeviceSweeperDataModel = (*customDeviceSweeperDataModel)(nil)

type (
	// DeviceSweeperDataModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDeviceSweeperDataModel.
	DeviceSweeperDataModel interface {
		deviceSweeperDataModel
		// Removed withSession as MongoDB handles sessions differently
	}

	customDeviceSweeperDataModel struct {
		*defaultDeviceSweeperDataModel
	}
)

// NewDeviceSweeperDataModel returns a model for the MongoDB collection.
func NewDeviceSweeperDataModel(client *mongo.Client) DeviceSweeperDataModel {
	return &customDeviceSweeperDataModel{
		defaultDeviceSweeperDataModel: newDeviceSweeperDataModel(client),
	}
}

