package model_test

import (
	"context"
	"testing"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/yourusername/yourproject/model"
)

type MockMongoCollection struct {
	mock.Mock
}

func (m *MockMongoCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func (m *MockMongoCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	args := m.Called(ctx, filter)
	return args.Get(0).(*mongo.SingleResult)
}

func (m *MockMongoCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).(*mongo.Cursor), args.Error(1)
}

func (m *MockMongoCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	args := m.Called(ctx, filter, update)
	return args.Get(0).(*mongo.UpdateResult), args.Error(1)
}

func (m *MockMongoCollection) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).(*mongo.DeleteResult), args.Error(1)
}

func TestDeviceCameraDataModel_Insert(t *testing.T) {
	ctx := context.TODO()
	mockCollection := new(MockMongoCollection)
	model := model.NewDeviceCameraDataModel(mockCollection)

	mockCollection.On("InsertOne", ctx, mock.Anything).Return(&mongo.InsertOneResult{}, nil)

	_, err := model.Insert(ctx, &model.DeviceCameraData{})
	assert.NoError(t, err)
	mockCollection.AssertExpectations(t)
}

func TestDeviceCameraDataModel_FindOne(t *testing.T) {
	ctx := context.TODO()
	mockCollection := new(MockMongoCollection)
	model := model.NewDeviceCameraDataModel(mockCollection)

	mockCollection.On("FindOne", ctx, bson.M{"_id": "123"}).Return(&mongo.SingleResult{}, nil)

	_, err := model.FindOne(ctx, "123")
	assert.NoError(t, err)
	mockCollection.AssertExpectations(t)
}

func TestDeviceCameraDataModel_FindOneByDeviceSn(t *testing.T) {
	ctx := context.TODO()
	mockCollection := new(MockMongoCollection)
	model := model.NewDeviceCameraDataModel(mockCollection)

	mockCollection.On("FindOne", ctx, bson.M{"deviceSn": "123"}).Return(&mongo.SingleResult{}, nil)

	_, err := model.FindOneByDeviceSn(ctx, "123")
	assert.NoError(t, err)
	mockCollection.AssertExpectations(t)
}

func TestDeviceCameraDataModel_FindByDeviceSnTimeRange(t *testing.T) {
	ctx := context.TODO()
	mockCollection := new(MockMongoCollection)
	model := model.NewDeviceCameraDataModel(mockCollection)

	mockCollection.On("Find", ctx, bson.M{"deviceSn": "123", "params.ct": bson.M{"$gte": "start", "$lte": "end"}}).Return(&mongo.Cursor{}, nil)

	_, err := model.FindByDeviceSnTimeRange(ctx, "123", "start", "end")
	assert.NoError(t, err)
	mockCollection.AssertExpectations(t)
}

func TestDeviceCameraDataModel_Update(t *testing.T) {
	ctx := context.TODO()
	mockCollection := new(MockMongoCollection)
	model := model.NewDeviceCameraDataModel(mockCollection)

	mockCollection.On("UpdateOne", ctx, bson.M{"_id": "123"}, bson.M{"$set": &model.DeviceCameraData{}}).Return(&mongo.UpdateResult{}, nil)

	_, err := model.Update(ctx, "123", &model.DeviceCameraData{})
	assert.NoError(t, err)
	mockCollection.AssertExpectations(t)
}

func TestDeviceCameraDataModel_Delete(t *testing.T) {
	ctx := context.TODO()
	mockCollection := new(MockMongoCollection)
	model := model.NewDeviceCameraDataModel(mockCollection)

	mockCollection.On("DeleteOne", ctx, bson.M{"_id": "123"}).Return(&mongo.DeleteResult{}, nil)

	_, err := model.Delete(ctx, "123")
	assert.NoError(t, err)
	mockCollection.AssertExpectations(t)
}

func TestDeviceCameraDataModel_Upsert(t *testing.T) {
	ctx := context.TODO()
	mockCollection := new(MockMongoCollection)
	model := model.NewDeviceCameraDataModel(mockCollection)

	mockCollection.On("UpdateOne", ctx, bson.M{"_id": "123"}, bson.M{"$set": &model.DeviceCameraData{}}, options.Update().SetUpsert(true)).Return(&mongo.UpdateResult{}, nil)

	data := []*model.DeviceCameraData{{ID: "123"}}
	err := model.Upsert(ctx, data)
	assert.NoError(t, err)
	mockCollection.AssertExpectations(t)
}