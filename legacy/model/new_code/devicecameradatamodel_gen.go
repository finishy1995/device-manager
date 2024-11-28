package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	DeviceCameraData struct {
		ID       string             `bson:"_id,omitempty"`
		DeviceSn string             `bson:"deviceSn"`
		Params   map[string]Param `bson:"params"`
	}

	Param struct {
		Pv string `bson:"pv"`
		Ct string `bson:"ct"`
		Ut string `bson:"ut"`
	}

	DeviceCameraDataModel struct {
		client *mongo.Client
		collection *mongo.Collection
	}
)

func NewDeviceCameraDataModel(client *mongo.Client) *DeviceCameraDataModel {
	collection := client.Database("test").Collection("device_camera_data")
	return &DeviceCameraDataModel{client: client, collection: collection}
}

func (m *DeviceCameraDataModel) Insert(ctx context.Context, data *DeviceCameraData) (*mongo.InsertOneResult, error) {
	return m.collection.InsertOne(ctx, data)
}

func (m *DeviceCameraDataModel) FindOne(ctx context.Context, id string) (*DeviceCameraData, error) {
	result := m.collection.FindOne(ctx, bson.M{"_id": id})
	var deviceData DeviceCameraData
	if err := result.Decode(&deviceData); err != nil {
		return nil, err
	}
	return &deviceData, nil
}

func (m *DeviceCameraDataModel) FindOneByDeviceSn(ctx context.Context, deviceSn string) (*DeviceCameraData, error) {
	result := m.collection.FindOne(ctx, bson.M{"deviceSn": deviceSn})
	var deviceData DeviceCameraData
	if err := result.Decode(&deviceData); err != nil {
		return nil, err
	}
	return &deviceData, nil
}

func (m *DeviceCameraDataModel) FindByDeviceSnTimeRange(ctx context.Context, deviceSn string, startTimestamp string, endTimestamp string) ([]*DeviceCameraData, error) {
	cursor, err := m.collection.Find(ctx, bson.M{"deviceSn": deviceSn, "params.ct": bson.M{"$gte": startTimestamp, "$lte": endTimestamp}})
	if err != nil {
		return nil, err
	}
	var results []*DeviceCameraData
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (m *DeviceCameraDataModel) Update(ctx context.Context, id string, data *DeviceCameraData) (*mongo.UpdateResult, error) {
	return m.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": data})
}

func (m *DeviceCameraDataModel) Delete(ctx context.Context, id string) (*mongo.DeleteResult, error) {
	return m.collection.DeleteOne(ctx, bson.M{"_id": id})
}

func (m *DeviceCameraDataModel) Upsert(ctx context.Context, data []*DeviceCameraData) error {
	for _, item := range data {
		_, err := m.collection.UpdateOne(ctx, bson.M{"_id": item.ID}, bson.M{"$set": item}, options.Update().SetUpsert(true))
		if err != nil {
			return err
		}
	}
	return nil
}
