package main

import (
	"fmt"
)

// Author:Boyn
// Date:2020/2/22

type UserInfo struct {
	Uid        int `PK`
	Username   string
	Departname string
	Created    string
}

func (u UserInfo) String() string {
	return fmt.Sprintf("%d %s %s %s", u.Uid, u.Username, u.Departname, u.Created)
}
