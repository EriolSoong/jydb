package row

type ID struct {
	ID string
}

func GenerateRowID() *ID {
	return &ID{}
}

func (id *ID) GetRowID(primaryKey string) string {
	return ""
}

func (id *ID) AutoIncrement() {
}

func (id *ID) Reset() {
}
