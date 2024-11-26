package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ DeviceCameraDataModel = (*customDeviceCameraDataModel)(nil)

type (
	// DeviceCameraDataModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDeviceCameraDataModel.
	DeviceCameraDataModel interface {
		deviceCameraDataModel
		withSession(session sqlx.Session) DeviceCameraDataModel
	}

	customDeviceCameraDataModel struct {
		*defaultDeviceCameraDataModel
	}
)

// NewDeviceCameraDataModel returns a model for the database table.
func NewDeviceCameraDataModel(conn sqlx.SqlConn) DeviceCameraDataModel {
	return &customDeviceCameraDataModel{
		defaultDeviceCameraDataModel: newDeviceCameraDataModel(conn),
	}
}

func (m *customDeviceCameraDataModel) withSession(session sqlx.Session) DeviceCameraDataModel {
	return NewDeviceCameraDataModel(sqlx.NewSqlConnFromSession(session))
}
