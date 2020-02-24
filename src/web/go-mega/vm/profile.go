package vm

import "web/go-mega/model"

// Author:Boyn
// Date:2020/2/24

// 个人主页VM
// 包含文章列表及其作者
type ProfileViewModel struct {
	BaseViewModel
	Posts       []model.Post
	ProfileUser model.User
	Editable    bool
}

type ProfileViewModelOp struct {
}

// 在这个函数中,包含了两个用户角色
// sUser是当前登录的用户
// pUser是我们准备要查看其个人主页的用户
func (ProfileViewModelOp) GetVM(sUser, pUser string) (ProfileViewModel, error) {
	v := ProfileViewModel{}
	v.SetTitle("Profile")
	u1, err := model.GetUserByUsername(pUser)
	if err != nil {
		return v, err
	}
	posts, _ := model.GetPostsByUserID(u1.ID)
	v.ProfileUser = *u1
	v.Posts = *posts
	v.Editable = sUser == pUser
	v.SetCurrentUser(sUser)
	return v, nil
}
