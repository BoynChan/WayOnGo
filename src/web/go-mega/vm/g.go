package vm

// Author:Boyn
// Date:2020/2/23

// 包含标题的VM
type BaseViewModel struct {
	Title       string
	CurrentUser string
	PrevPage    int // 上一页
	NextPage    int // 下一页
	Total       int // 总页数
	CurrentPage int // 当前页
	Limit       int // 每页限制的条目数
}

func (v *BaseViewModel) SetPrevAndNextPage() {
	if v.CurrentPage > 1 {
		v.PrevPage = v.CurrentPage - 1
	}
	if (v.Total-1)/v.Limit >= v.CurrentPage {
		v.NextPage = v.CurrentPage + 1
	}
}

func (v *BaseViewModel) SetBasePageViewModel(total, page, limit int) {
	v.Total = total
	v.CurrentPage = page
	v.Limit = limit
	v.SetPrevAndNextPage()
}

func (v *BaseViewModel) SetTitle(title string) {
	v.Title = title
}

func (v *BaseViewModel) SetCurrentUser(user string) {
	v.CurrentUser = user
}
