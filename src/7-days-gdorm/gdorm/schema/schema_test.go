package schema

import (
	"7-days-gdorm/gdorm/dialect"
	"fmt"
	"testing"
)

//Author: Boyn
//Date: 2020/3/25

type User struct {
	Name string `gdorm:"PRIMARY KEY"`
	Age  int
}

var TestDial, _ = dialect.GetDialect("sqlite3")

func TestParse(t *testing.T) {
	schema := Parse(&User{}, TestDial)
	if schema.Name != "User" || len(schema.Fields) != 2 {
		t.Fail()
	}
	fmt.Printf("%s %s\n", schema.Name, schema.FieldNames)
}
