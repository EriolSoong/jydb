package engine

//创建DB
func CreateDB(dbName string) error {

	return nil
}

//激活DB
func UseDB(dbName string) error {

	return nil
}

// 创建表
func CreateTable(table string, col []Column, primaryKey string) error {

	return nil
}

// 插入数据
func InsertData(table string, cols []Column) (bool, error) {

	return true, nil
}

// 插入多条数据
func InsertMultiData(table string, cols map[string][]Column) (int, error) {

	return 0, nil
}

// 查询数据
func GetLatestRecords(table string) (map[string][]Column, error) {

	return nil, nil
}

// 查询指定字段数据
func GetLatestRecordsByColNames(
	table string,
	colNames []string) (map[string][]Column, error) {

	return nil, nil
}

// 查询历史数据
func GetRecordsByTs(table string, timestamp int64) (map[string][]Column, error) {

	return nil, nil
}

// 根据字段查询历史数据
func GetRecordsByTsAndColNames(
	table string,
	primaryKey string,
	timestamp int64,
	colNames []string) (map[string][]Column, error) {

	return nil, nil
}

// 更新数据
func UpdateData(
	table string,
	primaryKey string,
	fields map[string][]Column) (successCount int, err error) {
	return 0, nil
}
