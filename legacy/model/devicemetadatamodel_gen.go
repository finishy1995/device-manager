// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.7.3

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	deviceMetadataFieldNames          = builder.RawFieldNames(&DeviceMetadata{})
	deviceMetadataRows                = strings.Join(deviceMetadataFieldNames, ",")
	deviceMetadataRowsExpectAutoSet   = strings.Join(stringx.Remove(deviceMetadataFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	deviceMetadataRowsWithPlaceHolder = strings.Join(stringx.Remove(deviceMetadataFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	deviceMetadataModel interface {
		Insert(ctx context.Context, data *DeviceMetadata) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*DeviceMetadata, error)
		FindByDeviceSn(ctx context.Context, deviceSn string) ([]*DeviceMetadata, error)
		FindOneByDeviceSnParamType(ctx context.Context, deviceSn string, paramType int64) (*DeviceMetadata, error)
		Update(ctx context.Context, data *DeviceMetadata) error
		Upsert(ctx context.Context, data []*DeviceMetadata) (*BatchResult, error)
		Delete(ctx context.Context, id int64) error
	}

	defaultDeviceMetadataModel struct {
		conn  sqlx.SqlConn
		table string
	}

	DeviceMetadata struct {
		Id         int64          `db:"id"`
		DeviceSn   string         `db:"device_sn"`   // device sn
		ParamType  int64          `db:"param_type"`  // param type
		ParamValue sql.NullString `db:"param_value"` // param value
		CreateTime time.Time      `db:"create_time"` // create time
		UpdateTime time.Time      `db:"update_time"` // update time
	}
)

func newDeviceMetadataModel(conn sqlx.SqlConn) *defaultDeviceMetadataModel {
	return &defaultDeviceMetadataModel{
		conn:  conn,
		table: "`device_metadata`",
	}
}

func (m *defaultDeviceMetadataModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultDeviceMetadataModel) FindOne(ctx context.Context, id int64) (*DeviceMetadata, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", deviceMetadataRows, m.table)
	var resp DeviceMetadata
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultDeviceMetadataModel) FindByDeviceSn(ctx context.Context, deviceSn string) ([]*DeviceMetadata, error) {
	query := fmt.Sprintf("select %s from %s where `device_sn` = ?", deviceMetadataRows, m.table)
	var resp []*DeviceMetadata
	err := m.conn.QueryRowsCtx(ctx, &resp, query, deviceSn)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultDeviceMetadataModel) FindOneByDeviceSnParamType(ctx context.Context, deviceSn string, paramType int64) (*DeviceMetadata, error) {
	var resp DeviceMetadata
	query := fmt.Sprintf("select %s from %s where `device_sn` = ? and `param_type` = ? limit 1", deviceMetadataRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, deviceSn, paramType)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultDeviceMetadataModel) Insert(ctx context.Context, data *DeviceMetadata) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, deviceMetadataRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.DeviceSn, data.ParamType, data.ParamValue)
	return ret, err
}

func (m *defaultDeviceMetadataModel) Update(ctx context.Context, newData *DeviceMetadata) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, deviceMetadataRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.DeviceSn, newData.ParamType, newData.ParamValue, newData.Id)
	return err
}

const BatchSize = 1000

// BatchResult stores the execution results of batch processing
type BatchResult struct {
    SuccessCount int      // Number of successfully processed records
    FailedBatch  []int    // Indexes of failed batches
    Err          error    // Last error encountered
}

func (m *defaultDeviceMetadataModel) Upsert(ctx context.Context, data []*DeviceMetadata) (*BatchResult, error) {
    if len(data) == 0 {
        return &BatchResult{}, nil
    }

    result := &BatchResult{}
    
    // Build base SQL statement
    baseSQL := fmt.Sprintf("insert into %s (%s) values ", m.table, deviceMetadataRowsExpectAutoSet)
    
    // Process data in batches
    for i := 0; i < len(data); i += BatchSize {
        end := i + BatchSize
        if end > len(data) {
            end = len(data)
        }
        
        // Get current batch data
        batchData := data[i:end]
        
        // Construct SQL statement
        var builder strings.Builder
        builder.WriteString(baseSQL)
        
        // Build parameter placeholders and values array
        values := make([]interface{}, 0, len(batchData)*3)
        for j := 0; j < len(batchData); j++ {
            if j > 0 {
                builder.WriteString(",")
            }
            builder.WriteString("(?,?,?)")
            values = append(values, batchData[j].DeviceSn, batchData[j].ParamType, batchData[j].ParamValue)
        }
        
        // Append on duplicate key update clause
        builder.WriteString(" on duplicate key update `param_value` = VALUES(`param_value`)")
        
        // Execute current batch within a transaction
        err := m.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
            _, err := session.Exec(builder.String(), values...)
            return err
        })
        
        if err != nil {
            result.FailedBatch = append(result.FailedBatch, i/BatchSize)
            result.Err = err
            return result, err
        }
        
        result.SuccessCount += len(batchData)
    }
    
    return result, nil
}

func (m *defaultDeviceMetadataModel) tableName() string {
	return m.table
}
