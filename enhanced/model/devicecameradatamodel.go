package model

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var _ DeviceCameraDataModel = (*customDeviceCameraDataModel)(nil)

type (
	// DeviceCameraDataModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDeviceCameraDataModel.
	DeviceCameraDataModel interface {
		deviceCameraDataModel
		// Removed withSession as MongoDB handles sessions differently
	}

	customDeviceCameraDataModel struct {
		*defaultDeviceCameraDataModel
	}
)



// NewDeviceCameraDataModel returns a model for the MongoDB collection.
func NewDeviceCameraDataModel(client *mongo.Client) DeviceCameraDataModel {
	return &customDeviceCameraDataModel{
		defaultDeviceCameraDataModel: newDeviceCameraDataModel(client),
	}
}

