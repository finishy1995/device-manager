package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ DeviceSweeperDataModel = (*customDeviceSweeperDataModel)(nil)

type (
	// DeviceSweeperDataModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDeviceSweeperDataModel.
	DeviceSweeperDataModel interface {
		deviceSweeperDataModel
		withSession(session sqlx.Session) DeviceSweeperDataModel
	}

	customDeviceSweeperDataModel struct {
		*defaultDeviceSweeperDataModel
	}
)

// NewDeviceSweeperDataModel returns a model for the database table.
func NewDeviceSweeperDataModel(conn sqlx.SqlConn) DeviceSweeperDataModel {
	return &customDeviceSweeperDataModel{
		defaultDeviceSweeperDataModel: newDeviceSweeperDataModel(conn),
	}
}

func (m *customDeviceSweeperDataModel) withSession(session sqlx.Session) DeviceSweeperDataModel {
	return NewDeviceSweeperDataModel(sqlx.NewSqlConnFromSession(session))
}
