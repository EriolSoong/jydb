package database

import (
	"jydb/gorocksdb"
	"jydb/storage/field"
	"jydb/storage/table"
	"time"
)

type Database struct {
	RDB        *gorocksdb.RocksDB
	Name       string
	UUID       string
	TableList  []*table.Table
	createTime time.Time
	modifyTime time.Time
}

// 创建数据库
// name 数据库保存路径
func CreateDatabase(name string) (*Database, error) {
	return nil, nil
}

// 打开已有数据库
// name 数据库所在位置
func OpenDatabase(name string) (*Database, error) {
	return nil, nil
}

// 重命名数据库
// newName 数据库名字
func (db *Database) Rename(newName string) error {
	return nil
}

// 删除数据库， 数据库标记为DELETED
func (db *Database) Drop() error {
	return nil
}

// 创建数据表
func (db *Database) CreateTable(name string, fields []*field.Field) (*table.Table, error) {
	return &table.Table{}, nil
}

// 获取数据表
func (db *Database) GetTable(name string) (*table.Table, error) {
	return &table.Table{}, nil
}
