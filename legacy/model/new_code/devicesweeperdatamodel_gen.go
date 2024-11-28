package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DeviceData struct {
	DeviceSn string             `bson:"deviceSn"`
	Params   map[string]ParamData `bson:"params"`
}

type ParamData struct {
	Pv string `bson:"pv"`
	Ct string `bson:"ct"`
	Ut string `bson:"ut"`
}

type DeviceDataRepository interface {
	Insert(ctx context.Context, data *DeviceData) error
	FindByDeviceSn(ctx context.Context, deviceSn string) (*DeviceData, error)
	Update(ctx context.Context, data *DeviceData) error
}

type defaultDeviceDataRepository struct {
	collection *mongo.Collection
}

func NewDeviceDataRepository(db *mongo.Database) DeviceDataRepository {
	return &defaultDeviceDataRepository{
		collection: db.Collection("deviceData"),
	}
}

func (r *defaultDeviceDataRepository) Insert(ctx context.Context, data *DeviceData) error {
	_, err := r.collection.InsertOne(ctx, data)
	return err
}

func (r *defaultDeviceDataRepository) FindByDeviceSn(ctx context.Context, deviceSn string) (*DeviceData, error) {
	filter := bson.M{"deviceSn": deviceSn}
	var result DeviceData
	err := r.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
}

func (r *defaultDeviceDataRepository) Update(ctx context.Context, data *DeviceData) error {
	filter := bson.M{"deviceSn": data.DeviceSn}
	update := bson.M{"$set": bson.M{"params": data.Params}}
	_, err := r.collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	return err
}
