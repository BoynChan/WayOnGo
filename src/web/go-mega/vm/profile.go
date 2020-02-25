package vm

import "web/go-mega/model"

// Author:Boyn
// Date:2020/2/24

// 个人主页VM
// 包含文章列表及其作者
type ProfileViewModel struct {
	BaseViewModel
	Posts          []model.Post // 文章列表
	ProfileUser    model.User   // 当前页面用户
	Editable       bool         // 是否可以编辑
	IsFollow       bool         // 是否被关注
	FollowersCount int          // 关注者数量
	FollowingCount int          // 关注的人数量

}

type ProfileViewModelOp struct {
}

// 在这个函数中,包含了两个用户角色
// sUser是当前登录的用户
// pUser是我们准备要查看其个人主页的用户
func (ProfileViewModelOp) GetVM(sUser, pUser string, page, limit int) (ProfileViewModel, error) {
	v := ProfileViewModel{}
	v.SetTitle("个人主页")
	p1, err := model.GetUserByUsername(pUser)
	if err != nil {
		return v, err
	}
	posts, total, _ := model.GetPostsByUserIDPageAndLimit(p1.ID, page, limit)
	v.ProfileUser = *p1
	v.Posts = *posts
	v.Editable = sUser == pUser
	if !v.Editable {
		// 先验证是否同一用户,如果是则不用进行是否关注的操作
		// 减少一次数据库的查询
		v.IsFollow = p1.IsFollowedByUser(sUser)
	}
	v.FollowersCount = p1.FollowersCount()
	v.FollowingCount = p1.FollowingCount()
	v.SetCurrentUser(sUser)
	v.SetBasePageViewModel(total, page, limit)
	return v, nil
}

// a关注了b
func Follow(a, b string) error {
	u, err := model.GetUserByUsername(a)
	if err != nil {
		return nil
	}
	return u.Follow(b)
}

// a取消关注了b
func UnFollow(a, b string) error {
	u, err := model.GetUserByUsername(a)
	if err != nil {
		return nil
	}
	return u.Unfollow(b)
}
