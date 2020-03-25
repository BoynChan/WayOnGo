package main

import (
	"7-days-gdorm/gdorm"
	"fmt"
	"testing"
)

//Author: Boyn
//Date: 2020/3/25

type User struct {
	Name string `gdorm:"PRIMARY KEY"`
	Age  int
}

func TestSession_CreateTable(t *testing.T) {
	engine, _ := gdorm.NewEngine("sqlite3", "/Users/chenzhanpeng/Code/Go/LearningGo/src/7-days-gdorm/gdorm/gee.db")
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
