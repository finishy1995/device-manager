package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ DeviceMetadataModel = (*customDeviceMetadataModel)(nil)

type (
	// DeviceMetadataModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDeviceMetadataModel.
	DeviceMetadataModel interface {
		deviceMetadataModel
		withSession(session sqlx.Session) DeviceMetadataModel
	}

	customDeviceMetadataModel struct {
		*defaultDeviceMetadataModel
	}
)

// NewDeviceMetadataModel returns a model for the database table.
func NewDeviceMetadataModel(conn sqlx.SqlConn) DeviceMetadataModel {
	return &customDeviceMetadataModel{
		defaultDeviceMetadataModel: newDeviceMetadataModel(conn),
	}
}

func (m *customDeviceMetadataModel) withSession(session sqlx.Session) DeviceMetadataModel {
	return NewDeviceMetadataModel(sqlx.NewSqlConnFromSession(session))
}
