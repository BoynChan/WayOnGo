package vm

import (
	"web/go-mega/model"
)

// Author:Boyn
// Date:2020/2/25
type EmailViewModel struct {
	Username string
	Token    string
	Server   string
}
type EmailViewModelOp struct {
}

func (EmailViewModelOp) GetVM(email string) EmailViewModel {
	v := EmailViewModel{}
	u, _ := model.GetUserByEmail(email)
	v.Username = u.Username
	v.Token, _ = u.GenerateToken()
	v.Server = "127.0.0.1"
	return v
}
