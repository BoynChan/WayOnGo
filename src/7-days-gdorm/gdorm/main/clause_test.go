package main

import (
	"7-days-gdorm/gdorm/clause"
	"reflect"
	"testing"
)

//Author: Boyn
//Date: 2020/3/26

func TestSelect(t *testing.T) {
	var c clause.Clause
	c.Set(clause.LIMIT, 3)
	c.Set(clause.SELECT, "User", []string{"*"})
	c.Set(clause.WHERE, "Name = ?", "Tom")
	c.Set(clause.ORDERBY, "Age ASC")
	sql, vars := c.Build(clause.SELECT, clause.WHERE, clause.ORDERBY, clause.LIMIT)
	t.Log(sql, vars)
	if sql != "SELECT * FROM User WHERE Name = ? ORDER BY Age ASC LIMIT ?" {
		t.Fail()
	}
	if !reflect.DeepEqual(vars, []interface{}{"Tom", 3}) {
		t.Fail()
	}
}
