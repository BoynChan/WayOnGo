package test

import (
	"7-days-gdorm/gdorm"
	"7-days-gdorm/gdorm/log"
	"7-days-gdorm/gdorm/session"
)

//Author: Boyn
//Date: 2020/3/26

type User struct {
	Name string `gdorm:"PRIMARY KEY"`
	Age  int
}

func getEngine() *gdorm.Engine {
	engine, _ := gdorm.NewEngine("sqlite3", "/Users/chenzhanpeng/Code/Go/LearningGo/src/7-days-gdorm/gdorm/gee.db")
	return engine
}

func (u *User) BeforeInsert(s *session.Session) error {
	log.Info("Before Insert", u)
	return nil
}

func (u *User) AfterQuery(s *session.Session) error {
	log.Info("After Query", u)
	return nil
}
