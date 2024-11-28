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

func TestInsert(t *testing.T) {
	mockCollection := new(MockCollection)
	repo := model.NewDeviceDataRepository(mockCollection)
	mockCollection.On("InsertOne", mock.Anything, mock.Anything).Return(&mongo.InsertOneResult{}, nil)
	deviceData := &model.DeviceData{DeviceSn: "123", Params: map[string]*model.DeviceParam{"param1": {Pv: "pv1", Ct: "ct1", Ut: "ut1"}}}
	err := repo.Insert(context.Background(), deviceData)
	assert.NoError(t, err)
	mockCollection.AssertExpectations(t)
}

func TestFindByDeviceSn(t *testing.T) {
	mockCollection := new(MockCollection)
	repo := model.NewDeviceDataRepository(mockCollection)
	mockCollection.On("FindOne", mock.Anything, bson.M{"deviceSn": "123"}).Return(&mongo.SingleResult{})
	_, err := repo.FindByDeviceSn(context.Background(), "123")
	assert.NoError(t, err)
	mockCollection.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	mockCollection := new(MockCollection)
	repo := model.NewDeviceDataRepository(mockCollection)
	deviceData := &model.DeviceData{DeviceSn: "123", Params: map[string]*model.DeviceParam{"param1": {Pv: "pv1", Ct: "ct1", Ut: "ut1"}}}
	mockCollection.On("UpdateOne", mock.Anything, bson.M{"deviceSn": "123"}, bson.M{"$set": bson.M{"params": deviceData.Params}}, options.Update().SetUpsert(true)).Return(&mongo.UpdateResult{}, nil)
	err := repo.Update(context.Background(), deviceData)
	assert.NoError(t, err)
	mockCollection.AssertExpectations(t)
}