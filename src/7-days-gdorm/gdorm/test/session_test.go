package test

import (
	"7-days-gdorm/gdorm/log"
	"7-days-gdorm/gdorm/session"
	"fmt"
	"testing"
)

//Author: Boyn
//Date: 2020/3/26

var (
	user1 = User{
		Name: "Jack",
		Age:  12,
	}
	user2 = User{
		Name: "Bob",
		Age:  13,
	}
	user3 = User{
		Name: "Rose",
		Age:  14,
	}
)

func testRecordInit(t *testing.T) *session.Session {
	engine := getEngine()
	s := engine.NewSession().Model(&User{})
	err1 := s.DropTable()
	err2 := s.CreateTable()
	_, err3 := s.Insert(user1, user2)
	if err1 != nil || err2 != nil || err3 != nil {
		t.Fail()
	}
	return s
}

func TestSessionInsert(t *testing.T) {
	s := testRecordInit(t)
	affected, err := s.Insert(user3)
	if err != nil || affected != 1 {
		t.Fail()
	}
}

func TestSessionFind(t *testing.T) {
	s := testRecordInit(t)
	users := make([]User, 0)
	if err := s.Find(&users); err != nil || len(users) != 2 {
		t.Fail()
	}
	fmt.Println(users)
}

func TestSessionUpdate(t *testing.T) {
	s := testRecordInit(t)
	affected, err := s.Where("Name = ?", "Jack").Update("Age", 11)
	if err != nil {
		log.Error(err)
		t.Fail()
	}
	log.Info(affected)
}

func TestSessionCount(t *testing.T) {
	s := testRecordInit(t)
	count, err := s.Count()
	if err != nil {
		log.Error(err)
		t.Fail()
	}
	fmt.Println(count)
}
