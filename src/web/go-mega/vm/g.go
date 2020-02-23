package vm

// Author:Boyn
// Date:2020/2/23

// 包含标题的VM
type BaseViewModel struct {
	Title string
}

func (v *BaseViewModel) SetTitle(title string) {
	v.Title = title
}
