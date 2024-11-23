package logic

import (
	"context"
	"database/sql"
	"finishy1995/device-manager/legacy/model"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var (
	// 任务队列
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

	// WaitGroup to wait for all goroutines
	var wg sync.WaitGroup
	wg.Add(int(t.thread))

	// Start specified number of goroutines
	for i := int32(0); i < t.thread; i++ {
		go func() {
			defer wg.Done()
			model := model.NewDeviceMetadataModel(sqlx.NewMysql(dataSource))

			// Run until context is done
			for {
				select {
				case <-ctx.Done():
					return
				default:
					data := generateDemoMetadataByRandomDevice(t.deviceNumber, t.deviceParamNumber, BatchSize)
					start := time.Now()
					result, err := model.Upsert(ctx, data)
					end := time.Now()
					if err != nil {
						fmt.Println("Upsert failed:", err)
						continue
					}
					t.latencyMicroseconds += end.Sub(start).Microseconds()
					t.successDeviceCount += int32(result.SuccessCount) / t.deviceParamNumber
				}
			}
		}()
	}

	// Wait for all goroutines to complete
	wg.Wait()
	t.success = true
}

func generateDemoMetadata(deviceNumber int32, deviceParamNumber int32) []*model.DeviceMetadata {
	result := make([]*model.DeviceMetadata, 0, deviceNumber)
	for i := 0; i < int(deviceNumber); i++ {
		for j := 0; j < int(deviceParamNumber); j++ {
			data := &model.DeviceMetadata{}
			mark := i%2 + 1
			data.DeviceSn = fmt.Sprintf("SN-%d%011d", mark, i)
			data.ParamType = int64(j)
			data.ParamValue = sql.NullString{String: fmt.Sprintf("%05d", rand.Intn(90000)+10000), Valid: true}
			result = append(result, data)
		}
	}
	return result
}

func generateDemoMetadataByRandomDevice(deviceNumber int32, deviceParamNumber int32, batchSize int) []*model.DeviceMetadata {
	result := make([]*model.DeviceMetadata, 0, batchSize)
	// Generate random device metadata batchSize times
	for i := 0; i < batchSize; i++ {
		for j := 0; j < int(deviceParamNumber); j++ {
			data := &model.DeviceMetadata{}

			// Generate random device number in range [0, deviceNumber)
			randomDevice := rand.Int31n(deviceNumber)
			mark := randomDevice%2 + 1
			// Generate device SN
			data.DeviceSn = fmt.Sprintf("SN-%d%011d", mark, randomDevice)
			data.ParamType = int64(j)
			data.ParamValue = sql.NullString{String: fmt.Sprintf("%05d", rand.Intn(90000)+10000), Valid: true}

			result = append(result, data)
		}
	}
	return result
}
