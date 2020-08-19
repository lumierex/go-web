package dialect

import (
	"fmt"
	"reflect"
	"time"
)

type mysql struct{}

// TODO (*mysql)(ni) 什么意思?
var _ Dialect = (*mysql)(nil)

func init() {
	RegisterDialect("mysql", &mysql{})
}

//  转换golang的类型到mysql中的类型
func (m *mysql) DataTypeOf(typ reflect.Value) string {
	switch typ.Kind() {
	case reflect.Bool:
		return "bool"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uintptr:
		return "int"
	case reflect.Int64:
		return "bigint"
	// 四个字节以及8个字节
	case reflect.Float32, reflect.Float64:
		return "float"
	case reflect.String:
		return "text"
	case reflect.Array, reflect.Slice:
		return "blob"
	case reflect.Struct:
		if _, ok := typ.Interface().(time.Time); ok {
			return "datetime"
		}
	}
	panic(fmt.Sprintf("invalid sql type %s (%s)", typ.Type().Name(), typ.Kind()))

}

// 返回表名是否存在的sql，以及表名参数
func (m *mysql) TableExistSQL(tableName string) (string, []interface{}) {
	args := []interface{}{tableName}
	return "select table_name from information_schema.tables where table_name =?", args
}
