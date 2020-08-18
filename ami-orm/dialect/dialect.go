package dialect

import "reflect"

// 不同数据库支持的数据库类型有差异

var dialectsMap = map[string]Dialect{}

type Dialect interface {
	// 将Go语言类型转为对应的数据库的类型
	DataTypeOf(typ reflect.Value) string
	// 返回某个表是否存在的SQL语句
	TableExistSQL(tableName string) (string, []interface{})
}

func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}

func GetDialect(name string) (Dialect, bool) {
	dialect, ok := dialectsMap[name]
	return dialect, ok
}
