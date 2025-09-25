// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemQueueMessageReceiveDao is the data access object for table system_queue_message_receive.
type SystemQueueMessageReceiveDao struct {
	table   string                           // table is the underlying table name of the DAO.
	group   string                           // group is the database configuration group name of current DAO.
	columns SystemQueueMessageReceiveColumns // columns contains all the column names of Table for convenient usage.
}

// SystemQueueMessageReceiveColumns defines and stores column names for table system_queue_message_receive.
type SystemQueueMessageReceiveColumns struct {
	MessageId  string // 队列消息主键
	UserId     string // 接收用户主键
	ReadStatus string // 已读状态 (1未读 2已读)
}

// systemQueueMessageReceiveColumns holds the columns for table system_queue_message_receive.
var systemQueueMessageReceiveColumns = SystemQueueMessageReceiveColumns{
	MessageId:  "message_id",
	UserId:     "user_id",
	ReadStatus: "read_status",
}

// NewSystemQueueMessageReceiveDao creates and returns a new DAO object for table data access.
func NewSystemQueueMessageReceiveDao() *SystemQueueMessageReceiveDao {
	return &SystemQueueMessageReceiveDao{
		group:   "default",
		table:   "system_queue_message_receive",
		columns: systemQueueMessageReceiveColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SystemQueueMessageReceiveDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SystemQueueMessageReceiveDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SystemQueueMessageReceiveDao) Columns() SystemQueueMessageReceiveColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SystemQueueMessageReceiveDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SystemQueueMessageReceiveDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SystemQueueMessageReceiveDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
