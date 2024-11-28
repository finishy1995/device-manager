package logic

import (
	"context"
	"finishy1995/device-manager/enhanced/model"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// Task queue
	taskList = map[int32]*task{}
)

type task struct {
	id                  int32
	success             bool
	deviceNumber        int32
	deviceParamNumber   int32
	thread              int32
	seconds             int32
	successDeviceCount  int32
	latencyMicroseconds int64
}

func NewTask(deviceNumber int32, deviceParamNumber int32, thread int32, seconds int32) *task {
	t := &task{
		id:                  rand.Int31(),
		deviceNumber:        deviceNumber,
		deviceParamNumber:   deviceParamNumber,
		thread:              thread,
		seconds:             seconds,
		successDeviceCount:  0,
		latencyMicroseconds: 0,
		success:             false,
	}
	taskList[t.id] = t
	return t
}

const BatchSize = 100

func (t *task) Run(dataSource string) {
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(t.seconds)*time.Second)
	defer cancel()

	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI(dataSource)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Failed to connect to MongoDB:", err)
		return
	}
	defer client.Disconnect(ctx)

	// WaitGroup to wait for all goroutines
	var wg sync.WaitGroup
	wg.Add(int(t.thread))

	// Calculate the range of devices each thread is responsible for
	segmentSize := t.deviceNumber / t.thread
	if t.deviceNumber%t.thread != 0 {
		segmentSize++ // Handle cases where devices cannot be evenly distributed
	}

	// Start specified number of goroutines
	for i := int32(0); i < t.thread; i++ {
		startDevice := i * segmentSize
		endDevice := (i+1)*segmentSize - 1
		if endDevice >= t.deviceNumber {
			endDevice = t.deviceNumber - 1 // Ensure the last segment doesn't exceed total devices
		}
		go func(start, end int32) {
			defer wg.Done()
			model := model.NewDeviceMetadataModel(client)

			// Run until context is done
			for {
				select {
				case <-ctx.Done():
					return
				default:
					data := generateDemoMetadataByDeviceRange(start, end, t.deviceParamNumber, BatchSize)
					start := time.Now()
					//result, err := model.Upsert(ctx, data)
					_, err := model.Upsert(ctx, data)
					end := time.Now()
					if err != nil {
						fmt.Println("Upsert failed:", err)
						continue
					}
					t.latencyMicroseconds += end.Sub(start).Microseconds()
					//t.successDeviceCount += int32(result.ModifiedCount) / t.deviceParamNumber
				}
			}
		}(startDevice, endDevice)
	}

	// Wait for all goroutines to complete
	wg.Wait()
	t.success = true
}

func generateDemoMetadata(deviceNumber int32, deviceParamNumber int32) []*model.DeviceMetadata {
	result := make([]*model.DeviceMetadata, 0, deviceNumber)
	for i := 0; i < int(deviceNumber); i++ {
		device := &model.DeviceMetadata{
			DeviceSn: fmt.Sprintf("SN-%d%011d", i%2+1, i),
			Params:   make(map[string]model.Param),
		}
		for j := 0; j < int(deviceParamNumber); j++ {
			paramType := fmt.Sprintf("%d", j)
			paramValue := fmt.Sprintf("%05d", rand.Intn(90000)+10000)
			currentTime := time.Now()
			device.Params[paramType] = model.Param{
				PV: paramValue,
				CT: currentTime,
				UT: currentTime,
			}
		}
		result = append(result, device)
	}
	return result
}

func generateDemoMetadataByDeviceRange(startDevice, endDevice int32, deviceParamNumber int32, batchSize int) []*model.DeviceMetadata {
	result := make([]*model.DeviceMetadata, 0, batchSize)
	// Generate random device metadata within the specified range
	for i := 0; i < batchSize; i++ {
		// Select a random device number within the range [startDevice, endDevice]
		randomDevice := rand.Int31n(endDevice-startDevice+1) + startDevice
		// Create a new DeviceMetadata instance
		device := &model.DeviceMetadata{
			DeviceSn: fmt.Sprintf("SN-%d%011d", randomDevice%2+1, randomDevice),
			Params:   make(map[string]model.Param),
		}
		// Generate parameters for the device
		for j := 0; j < int(deviceParamNumber); j++ {
			paramType := fmt.Sprintf("%d", j)
			paramValue := fmt.Sprintf("%05d", rand.Intn(90000)+10000)
			currentTime := time.Now()
			device.Params[paramType] = model.Param{
				PV: paramValue,
				CT: currentTime,
				UT: currentTime,
			}
		}
		result = append(result, device)
	}
	return result
}

