package session

import (
	"7-days-gdorm/gdorm/clause"
	"reflect"
)

//Author: Boyn
//Date: 2020/3/26
// this file is to record crud code

// INSERT INTO tableName( cols... ) VALUES ( value1 ), ( value2 ), ( value3 ) ...
// this function is to convert values into right sentence
// first, we will convert values into recordValues, which is the sorted field name set
// we should make sure that every values is in the same type of struct or will cause unknown error
// second, we build a sql sentence which contain the first sentence in this section but in a different way(we can see it in log output).
// at last, we will exec this sentence by session.Raw() and return the rows that inserted.
func (s *Session) Insert(values ...interface{}) (int64, error) {
	recordValues := make([]interface{}, 0)
	for _, value := range values {
		table := s.Model(value).RefTable()
		s.clause.Set(clause.INSERT, table.Name, table.FieldNames)
		recordValues = append(recordValues, table.RecordValues(value))
	}
	s.clause.Set(clause.VALUES, recordValues...)
	sql, vars := s.clause.Build(clause.INSERT, clause.VALUES)
	result, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// select *  from $tableName where $field1=$value1 and $field2=$value2
// find function is to find all record and put it into values as a slice
func (s *Session) Find(values interface{}) error {
	destSlice := reflect.Indirect(reflect.ValueOf(values))
	destType := destSlice.Type().Elem()
	table := s.Model(reflect.New(destType).Elem().Interface()).RefTable()

	s.clause.Set(clause.SELECT, table.Name, table.FieldNames)
	sql, vars := s.clause.Build(clause.SELECT, clause.WHERE, clause.ORDERBY, clause.LIMIT)
	rows, err := s.Raw(sql, vars...).QueryRows()
	if err != nil {
		return err
	}

	for rows.Next() {
		dest := reflect.New(destType).Elem()
		var values []interface{}
		for _, name := range table.FieldNames {
			values = append(values, dest.FieldByName(name).Addr().Interface())
		}
		if err := rows.Scan(values...); err != nil {
			return err
		}
		destSlice.Set(reflect.Append(destSlice, dest))
	}
	return rows.Close()
}
