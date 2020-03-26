package session

import (
	"7-days-gdorm/gdorm/clause"
	"7-days-gdorm/gdorm/log"
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
	log.Info(sql, vars)
	result, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
