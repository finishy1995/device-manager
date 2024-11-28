package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	DeviceCameraData struct {
		Id           string                 `bson:"_id"`
		DeviceSn     string                 `bson:"deviceSn"`
		Params       map[string]interface{} `bson:"params"`
	}

	DeviceCameraDataModel interface {
		Insert(ctx context.Context, data *DeviceCameraData) (*mongo.InsertOneResult, error)
		FindOne(ctx context.Context, id string) (*DeviceCameraData, error)
		FindOneByDeviceSn(ctx context.Context, deviceSn string) (*DeviceCameraData, error)
		Update(ctx context.Context, data *DeviceCameraData) (*mongo.UpdateResult, error)
		Delete(ctx context.Context, id string) (*mongo.DeleteResult, error)
	}

	defaultDeviceCameraDataModel struct {
		client *mongo.Client
		collection *mongo.Collection
	}
)

func NewDeviceCameraDataModel(client *mongo.Client) *defaultDeviceCameraDataModel {
	return &defaultDeviceCameraDataModel{
		client: client,
		collection: client.Database("test").Collection("device_camera_data"),
	}
}

func (m *defaultDeviceCameraDataModel) Insert(ctx context.Context, data *DeviceCameraData) (*mongo.InsertOneResult, error) {
	return m.collection.InsertOne(ctx, data)
}

func (m *defaultDeviceCameraDataModel) FindOne(ctx context.Context, id string) (*DeviceCameraData, error) {
	var result DeviceCameraData
	filter := bson.D{{"_id", id}}
	err := m.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (m *defaultDeviceCameraDataModel) FindOneByDeviceSn(ctx context.Context, deviceSn string) (*DeviceCameraData, error) {
	var result DeviceCameraData
	filter := bson.D{{"deviceSn", deviceSn}}
	err := m.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (m *defaultDeviceCameraDataModel) Update(ctx context.Context, data *DeviceCameraData) (*mongo.UpdateResult, error) {
	filter := bson.D{{"_id", data.Id}}
	update := bson.D{{"$set", bson.D{{"deviceSn", data.DeviceSn}, {"params", data.Params}}}}
	return m.collection.UpdateOne(ctx, filter, update)
}

func (m *defaultDeviceCameraDataModel) Delete(ctx context.Context, id string) (*mongo.DeleteResult, error) {
	filter := bson.D{{"_id", id}}
	return m.collection.DeleteOne(ctx, filter)
}
