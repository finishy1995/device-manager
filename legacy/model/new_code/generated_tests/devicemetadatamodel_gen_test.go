package model_test

import (
	"context"
	"testing"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/your_project/model"
)

type MockCollection struct {
	mock.Mock
}

func (m *MockCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func (m *MockCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	args := m.Called(ctx, filter)
	return args.Get(0).(*mongo.SingleResult)
}

func (m *MockCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	args := m.Called(ctx, filter, update)
	return args.Get(0).(*mongo.UpdateResult), args.Error(1)
}

func (m *MockCollection) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).(*mongo.DeleteResult), args.Error(1)
}

func TestInsert(t *testing.T) {
	mockCollection := new(MockCollection)
	repo := model.NewDeviceDataRepository(mockCollection)
	deviceData := &model.DeviceData{Id: "123", DeviceSn: "abc", Params: map[string]model.Param{"param1": {Pv: "pv1", Ct: "ct1", Ut: "ut1"}}}

	mockCollection.On("InsertOne", mock.Anything, deviceData).Return(nil, nil)

	err := repo.Insert(context.Background(), deviceData)

	assert.NoError(t, err)
	mockCollection.AssertExpectations(t)
}

func TestFindOne(t *testing.T) {
	mockCollection := new(MockCollection)
	repo := model.NewDeviceDataRepository(mockCollection)
	deviceData := &model.DeviceData{Id: "123", DeviceSn: "abc", Params: map[string]model.Param{"param1": {Pv: "pv1", Ct: "ct1", Ut: "ut1"}}}

	mockCollection.On("FindOne", mock.Anything, bson.M{"_id": "123"}).Return(deviceData, nil)

	result, err := repo.FindOne(context.Background(), "123")

	assert.NoError(t, err)
	assert.Equal(t, deviceData, result)
	mockCollection.AssertExpectations(t)
}

func TestFindByDeviceSn(t *testing.T) {
	mockCollection := new(MockCollection)
	repo := model.NewDeviceDataRepository(mockCollection)
	deviceData := &model.DeviceData{Id: "123", DeviceSn: "abc", Params: map[string]model.Param{"param1": {Pv: "pv1", Ct: "ct1", Ut: "ut1"}}}

	mockCollection.On("FindOne", mock.Anything, bson.M{"deviceSn": "abc"}).Return(deviceData, nil)

	result, err := repo.FindByDeviceSn(context.Background(), "abc")

	assert.NoError(t, err)
	assert.Equal(t, deviceData, result)
	mockCollection.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	mockCollection := new(MockCollection)
	repo := model.NewDeviceDataRepository(mockCollection)
	deviceData := &model.DeviceData{Id: "123", DeviceSn: "abc", Params: map[string]model.Param{"param1": {Pv: "pv1", Ct: "ct1", Ut: "ut1"}}}

	mockCollection.On("UpdateOne", mock.Anything, bson.M{"_id": "123"}, bson.M{"$set": deviceData}).Return(nil, nil)

	err := repo.Update(context.Background(), deviceData)

	assert.NoError(t, err)
	mockCollection.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	mockCollection := new(MockCollection)
	repo := model.NewDeviceDataRepository(mockCollection)

	mockCollection.On("DeleteOne", mock.Anything, bson.M{"_id": "123"}).Return(nil, nil)

	err := repo.Delete(context.Background(), "123")

	assert.NoError(t, err)
	mockCollection.AssertExpectations(t)
}