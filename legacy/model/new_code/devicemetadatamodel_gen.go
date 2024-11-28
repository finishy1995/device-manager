package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	DeviceData struct {
		Id     string            `bson:"_id"`
		DeviceSn string          `bson:"deviceSn"`
		Params  map[string]Param `bson:"params"`
	}

	Param struct {
		Pv string `bson:"pv"`
		Ct string `bson:"ct"`
		Ut string `bson:"ut"`
	}

	DeviceDataRepository interface {
		Insert(ctx context.Context, data *DeviceData) error
		FindOne(ctx context.Context, id string) (*DeviceData, error)
		FindByDeviceSn(ctx context.Context, deviceSn string) (*DeviceData, error)
		Update(ctx context.Context, data *DeviceData) error
		Delete(ctx context.Context, id string) error
	}

	defaultDeviceDataRepository struct {
		collection *mongo.Collection
	}
)

func NewDeviceDataRepository(db *mongo.Database) DeviceDataRepository {
	return &defaultDeviceDataRepository{
		collection: db.Collection("deviceData"),
	}
}

func (r *defaultDeviceDataRepository) Insert(ctx context.Context, data *DeviceData) error {
	_, err := r.collection.InsertOne(ctx, data)
	return err
}

func (r *defaultDeviceDataRepository) FindOne(ctx context.Context, id string) (*DeviceData, error) {
	var result DeviceData
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *defaultDeviceDataRepository) FindByDeviceSn(ctx context.Context, deviceSn string) (*DeviceData, error) {
	var result DeviceData
	err := r.collection.FindOne(ctx, bson.M{"deviceSn": deviceSn}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *defaultDeviceDataRepository) Update(ctx context.Context, data *DeviceData) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": data.Id}, bson.M{"$set": data})
	return err
}

func (r *defaultDeviceDataRepository) Delete(ctx context.Context, id string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
