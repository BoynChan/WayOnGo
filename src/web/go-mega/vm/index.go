package vm

import (
	"web/go-mega/model"
)

// Author:Boyn
// Date:2020/2/23

// 包含标题与文章的VM
type IndexViewModel struct {
	User model.User
	BaseViewModel
	Posts []model.Post
}

// IndexViewModel的操作结构
type IndexViewModelOp struct {
}

func (i IndexViewModelOp) GetVM() IndexViewModel {
	username, _ := model.GetUserByUsername("rene")
	posts, _ := model.GetPostsByUserID(username.ID)
	v := IndexViewModel{
		User:          *username,
		BaseViewModel: BaseViewModel{Title: "HomePage"},
		Posts:         *posts,
	}
	return v
}
