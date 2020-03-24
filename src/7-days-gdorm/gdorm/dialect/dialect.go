package dialect

import "reflect"

// Author:Boyn
// Date:2020/3/23
// get the same interface and operate from different db
var dialectsMap = map[string]Dialect{}

type Dialect interface {
	// transform go type into db data type.
	DataTypeOf(typ reflect.Value) string
	// return if specified table exist.
	TableExistSQL(tableName string) (string, []interface{})
}

func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}

func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name]
	return
}
