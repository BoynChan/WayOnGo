package test

import (
	"7-days-gdorm/gdorm/dialect"
	schema2 "7-days-gdorm/gdorm/schema"
	"fmt"
	"testing"
)

//Author: Boyn
//Date: 2020/3/25

var TestDial, _ = dialect.GetDialect("sqlite3")

func TestParse(t *testing.T) {
	schema := schema2.Parse(&User{}, TestDial)
	if schema.Name != "User" || len(schema.Fields) != 2 {
		t.Fail()
	}
	fmt.Printf("%s %s\n", schema.Name, schema.FieldNames)
}
