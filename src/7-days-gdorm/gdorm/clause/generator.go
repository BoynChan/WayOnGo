package clause

import (
	"fmt"
	"strings"
)

//Author: Boyn
//Date: 2020/3/25

//define generator function which can be used to generator specific SQL sentence
type generator func(values ...interface{}) (string, []interface{})

// define generators map to store different type of generators
// the key Type defined in clause.go which represent different SQL sentence type like insert select etc...
var generators map[Type]generator

func init() {
	generators = make(map[Type]generator)
	generators[INSERT] = _insert
	generators[VALUES] = _values
	generators[SELECT] = _select
	generators[LIMIT] = _limit
	generators[WHERE] = _where
	generators[ORDERBY] = _orderBy
}

// mark the number of parameter and add ? to sentence
func genBindVars(num int) string {
	var vars []string
	for i := 0; i < num; i++ {
		vars = append(vars, "?")
	}
	return strings.Join(vars, ", ")
}

// INSERT INTO $tableName ( $fields...)
// we use values[0] as tableName and values[1] as fields (which is a slice of string)
func _insert(values ...interface{}) (string, []interface{}) {
	tableName := values[0]
	fields := strings.Join(values[1].([]string), ",")
	return fmt.Sprintf("INSERT INTO %s (%v)", tableName, fields), []interface{}{}
}

// generate the sentence like VALUES ($1),($2),(...)
func _values(values ...interface{}) (string, []interface{}) {
	var bindStr string
	var sql strings.Builder
	var vars []interface{}

	sql.WriteString("VALUES ")
	for i, value := range values {
		v := value.([]interface{})
		if bindStr == "" {
			bindStr = genBindVars(len(v))
		}
		sql.WriteString(fmt.Sprintf("(%v)", bindStr))
		if i+1 != len(values) {
			sql.WriteString(", ")
		}
		vars = append(vars, v...)
	}
	return sql.String(), vars
}

// SELECT $fields FROM $tableName
// we use $0 as tableNmae
// and $1 will be a string slice as fields
func _select(values ...interface{}) (string, []interface{}) {
	tableName := values[0]
	fields := strings.Join(values[1].([]string), ",")
	return fmt.Sprintf("SELECT %v FROM %s", fields, tableName), []interface{}{}
}

// LIMIT $value
func _limit(values ...interface{}) (string, []interface{}) {
	return "LIMIT ?", values
}

// WHERE $desc
func _where(values ...interface{}) (string, []interface{}) {
	desc, vars := values[0], values[1:]
	return fmt.Sprintf("WHERE %s", desc), vars
}

// ORDER BY $1,$2...
func _orderBy(values ...interface{}) (string, []interface{}) {
	return fmt.Sprintf("ORDER BY %s", values[0]), []interface{}{}
}
