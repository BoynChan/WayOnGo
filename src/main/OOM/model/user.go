package model

import "fmt"

//Author: Boyn
//Date: 2020/3/29

type User struct {
	Username string
	Password string
	Age      int
}

func NewUser(name, password string, age int) *User {
	return &User{
		Username: name,
		Password: password,
		Age:      age,
	}
}

func (u *User) SetUsername(username string) {
	u.logForChange("Username", username)
	u.Username = username
}

func (u *User) SetAge(i int) {
	u.logForChange("Age", i)
	u.Age = i
}

func (u *User) SetPassword(password string) {
	u.logForChange("Password", password)
	u.Password = password
}

func (u *User) logForChange(field string, value interface{}) {
	fmt.Printf("%s changed to %v", field, value)
}
