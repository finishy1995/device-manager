package model

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var _ DeviceMetadataModel = (*customDeviceMetadataModel)(nil)

type (
	// DeviceMetadataModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDeviceMetadataModel.
	DeviceMetadataModel interface {
		deviceMetadataModel
		// Removed withSession as MongoDB handles sessions differently
	}

	customDeviceMetadataModel struct {
		*defaultDeviceMetadataModel
	}
)

// NewDeviceMetadataModel returns a model for the MongoDB collection.
func NewDeviceMetadataModel(client *mongo.Client) DeviceMetadataModel {
	return &customDeviceMetadataModel{
		defaultDeviceMetadataModel: newDeviceMetadataModel(client),
	}
}

