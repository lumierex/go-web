package dialect

import "reflect"

type mysql struct{}

// TODO (*mysql)(ni) 什么意思?
var _ Dialect = (*mysql)(nil)

func init() {
	RegisterDialect("mysql", &mysql{})
}

func (m *mysql) DataTypeOf(typ reflect.Value) string {
	return ""
}

func (m *mysql) TableExistSQL(tableName string) (string, []interface{}) {
	args := []interface{}{tableName}
	//return "SELECT name "
	return "", args
}
