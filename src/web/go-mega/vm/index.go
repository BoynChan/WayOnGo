package vm

import (
	"web/go-mega/model"
)

// Author:Boyn
// Date:2020/2/23

// 包含标题与文章的VM
type IndexViewModel struct {
	BaseViewModel
	Posts []model.Post
	Flash string
}

// IndexViewModel的操作结构
type IndexViewModelOp struct {
}

func (i IndexViewModelOp) GetVM(username string, flash string, page, limit int) IndexViewModel {
	user, _ := model.GetUserByUsername(username)
	// 获取关注者的文章
	posts, total, _ := user.FollowingPostsByPageAndLimit(page, limit)
	v := IndexViewModel{BaseViewModel{}, *posts, flash}
	v.SetTitle("HomePage")
	v.SetCurrentUser(username)
	v.SetBasePageViewModel(total, page, limit)
	return v
}

func CreatePost(username, post string) error {
	u, _ := model.GetUserByUsername(username)
	return u.CreatePost(post)
}
