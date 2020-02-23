package vm

import (
	. "web/go-mega/model"
)

// Author:Boyn
// Date:2020/2/23

// 包含标题与文章的VM
type IndexViewModel struct {
	User User
	BaseViewModel
	Posts []Post
}

// IndexViewModel的操作结构
type IndexViewModelOp struct {
}

func (i IndexViewModelOp) GetVM() IndexViewModel {
	u1 := User{Username: "boyn"}
	u2 := User{Username: "kaserchan"}
	posts := []Post{
		{
			User:  u1,
			Title: "模板的if-else操作",
		},
		{
			User:  u1,
			Title: "模板的插值操作",
		},
		{
			User:  u2,
			Title: "模板的循环操作",
		},
	}
	v := IndexViewModel{
		User:          u1,
		BaseViewModel: BaseViewModel{Title: "HomePage"},
		Posts:         posts,
	}
	return v
}
