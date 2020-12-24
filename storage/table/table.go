package table

import (
	"jydb/storage/field"
	"jydb/storage/row"
	"time"
)

type Table struct {
	Name       string
	DbUUId     string
	UUID       string
	RowID      *row.ID
	createTime time.Time
	modifyTime time.Time
	Fields     map[string][]*field.Field //map是rowID:fields
	PKToID     map[string]string         //主键到行号的索引
	IDToPK     map[string]string         //行号到主键的索引
	UKToID     map[string]string         //唯一键到行号的索引
}

// 重命名数据表
func (db *Table) Rename(newName string) error {
	return nil
}

// 插入一条数据，rowID自动生成
func (db *Table) InsertRow(fields []*field.Field) error {
	return nil
}

// 插入多条数据，rowID自动生成
func (db *Table) InsertRows(rowToFields map[string][]*field.Field) error {
	return nil
}

//根据时间戳查询结果 t=nil 查询最新结果
//return: rowID:[]*Field
func (db *Table) GetAll(t *time.Time) (map[string][]*field.Field, error) {
	return map[string][]*field.Field{}, nil
}

//根据时间查询指定rowID的结果
func (db *Table) GetRow(rowID string, t *time.Time) ([]*field.Field, error) {
	return []*field.Field{}, nil
}

//根据时间查询指定主键的结果
func (db *Table) GetRowByPK(PK string, t *time.Time) ([]*field.Field, error) {
	return []*field.Field{}, nil
}

//根据时间查询指定唯一键的结果
func (db *Table) GetRowByUK(UK string, t *time.Time) ([]*field.Field, error) {
	return []*field.Field{}, nil
}

//根据列名查询该列所有数据
func (db *Table) GetFields(fieldNames []string, t *time.Time) (map[string][]*field.Field, error) {
	return map[string][]*field.Field{}, nil
}

//并不是真的删除字段，只是添加DELETED标识
func (db *Table) DropRow(rowID string) error {
	return nil
}

//获取主键
func (db *Table) GetPK() (*field.Field, error) {
	return &field.Field{}, nil
}

//获取唯一键
func (db *Table) GetUK() ([]*field.Field, error) {
	return []*field.Field{}, nil
}

//并不是真的删除，只是添加DELETED标识
func (db *Table) Drop() error {
	return nil
}
