package engine

import (
	"jydb/storage/database"
	"jydb/storage/table"
)

type DBEngine struct {
	dbs          map[string]*database.Database //数据库连接池，保持长连接
	activeDB     *database.Database
	activeTables map[string]*table.Table //当前操作的数据表
}

func (e *DBEngine) init() error {
	return nil
}

func (e *DBEngine) destroy() {

}

var Engine *DBEngine

func Run() {
	Engine = new(DBEngine)
	err := Engine.init()
	if err != nil {
		panic(err)
	}
}
