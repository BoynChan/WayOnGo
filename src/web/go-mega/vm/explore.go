package vm

import "web/go-mega/model"

// Author:Boyn
// Date:2020/2/25

type ExploreViewModel struct {
	BaseViewModel
	Posts *[]model.Post
}

type ExploreViewModelOp struct {
}

func (ExploreViewModelOp) GetVM(page, limit int) (ExploreViewModel, error) {
	v := ExploreViewModel{}
	posts, total, _ := model.GetPostsByPageAndLimit(page, limit)
	v.Posts = posts
	v.Title = "探索"
	v.SetBasePageViewModel(total, page, limit)
	return v, nil
}
