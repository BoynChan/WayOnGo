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

func (i IndexViewModelOp) GetVM(username string) IndexViewModel {
	user, _ := model.GetUserByUsername(username)
	posts, _ := model.GetPostsByUserID(user.ID)
	v := IndexViewModel{
		User:          *user,
		BaseViewModel: BaseViewModel{Title: "HomePage", CurrentUser: username},
		Posts:         *posts,
	}
	return v
}
