package test

import (
	"fmt"
	"testing"
)

//Author: Boyn
//Date: 2020/3/25

func TestSession_CreateTable(t *testing.T) {
	engine := getEngine()
	defer engine.Close()
	s := engine.NewSession()
	s.Model(&User{})
	err := s.DropTable()
	if err != nil {
		fmt.Println(err)
	}
	err = s.CreateTable()
	if err != nil {
		fmt.Println(err)
	}
	if !s.HasTable() {
		t.Fail()
	}
}
