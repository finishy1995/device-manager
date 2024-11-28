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

type MockMongoClient struct {
	mock.Mock
}

func (m *MockMongoClient) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func (m *MockMongoClient) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	args := m.Called(ctx, filter)
	return args.Get(0).(*mongo.SingleResult)
}

func (m *MockMongoClient) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	args := m.Called(ctx, filter, update)
	return args.Get(0).(*mongo.UpdateResult), args.Error(1)
}

func (m *MockMongoClient) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).(*mongo.DeleteResult), args.Error(1)
}

func TestInsert(t *testing.T) {
	mockClient := new(MockMongoClient)
	model := model.NewDeviceCameraDataModel(mockClient)

	mockClient.On("InsertOne", mock.Anything, mock.Anything).Return(&mongo.InsertOneResult{}, nil)

	_, err := model.Insert(context.Background(), &model.DeviceCameraData{})
	assert.NoError(t, err)
}

func TestFindOne(t *testing.T) {
	mockClient := new(MockMongoClient)
	model := model.NewDeviceCameraDataModel(mockClient)

	mockClient.On("FindOne", mock.Anything, mock.Anything).Return(&mongo.SingleResult{})

	_, err := model.FindOne(context.Background(), "123")
	assert.NoError(t, err)
}

func TestFindOneByDeviceSn(t *testing.T) {
	mockClient := new(MockMongoClient)
	model := model.NewDeviceCameraDataModel(mockClient)

	mockClient.On("FindOne", mock.Anything, mock.Anything).Return(&mongo.SingleResult{})

	_, err := model.FindOneByDeviceSn(context.Background(), "123")
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	mockClient := new(MockMongoClient)
	model := model.NewDeviceCameraDataModel(mockClient)

	mockClient.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).Return(&mongo.UpdateResult{}, nil)

	_, err := model.Update(context.Background(), &model.DeviceCameraData{})
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	mockClient := new(MockMongoClient)
	model := model.NewDeviceCameraDataModel(mockClient)

	mockClient.On("DeleteOne", mock.Anything, mock.Anything).Return(&mongo.DeleteResult{}, nil)

	_, err := model.Delete(context.Background(), "123")
	assert.NoError(t, err)
}