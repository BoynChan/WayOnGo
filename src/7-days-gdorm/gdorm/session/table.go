package session

import (
	"7-days-gdorm/gdorm/log"
	"7-days-gdorm/gdorm/schema"
	"fmt"
	"reflect"
	"strings"
)

//Author: Boyn
//Date: 2020/3/25

// this function can parse a value into its table model
// if s.refTable is nil or we use a different type of value as parameter, session s will resolve this value
// and store corresponding table of value into s.refTable
func (s *Session) Model(value interface{}) *Session {
	if s.refTable == nil || reflect.TypeOf(value) != reflect.TypeOf(s.refTable.Model) {
		s.refTable = schema.Parse(value, s.dialect)
	}
	return s
}

func (s *Session) RefTable() *schema.Schema {
	if s.refTable == nil {
		log.Error("Model is not set")
	}
	return s.refTable
}

func (s *Session) CreateTable() error {
	table := s.RefTable()
	var columns []string
	for _, field := range table.Fields {
		columns = append(columns, fmt.Sprintf("%s %s %s", field.Name, field.Type, field.Tag))
	}
	desc := strings.Join(columns, ",")
	_, err := s.Raw(fmt.Sprintf("create table %s (%s)", table.Name, desc)).Exec()
	return err
}

func (s *Session) DropTable() error {
	_, err := s.Raw(fmt.Sprintf("drop table if exists %s;", s.RefTable().Name)).Exec()
	return err
}

func (s *Session) HasTable() bool {
	sql, values := s.dialect.TableExistSQL(s.RefTable().Name)
	row := s.Raw(sql, values...).QueryRow()
	var tmp string
	err := row.Scan(&tmp)
	if err != nil {
		log.Error(err)
		return false
	}
	return tmp == s.RefTable().Name
}
