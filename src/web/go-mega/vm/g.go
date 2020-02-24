package vm

// Author:Boyn
// Date:2020/2/23

// 包含标题的VM
type BaseViewModel struct {
	Title       string
	CurrentUser string
}

func (v *BaseViewModel) SetTitle(title string) {
	v.Title = title
}

func (v *BaseViewModel) SetCurrentUser(user string) {
	v.CurrentUser = user
}
