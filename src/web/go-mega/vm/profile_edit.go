package vm

import "web/go-mega/model"

// Author:Boyn
// Date:2020/2/24

type ProfileEditViewModel struct {
	LoginViewModel
	ProfileUser model.User
}

type ProfileEditViewModelOP struct {
}

func (ProfileEditViewModelOP) GetVM(username string) ProfileEditViewModel {
	v := ProfileEditViewModel{}
	u, _ := model.GetUserByUsername(username)
	v.SetTitle("Profile Edit")
	v.SetCurrentUser(username)
	v.ProfileUser = *u
	return v
}

func UpdateAboutMe(username, text string) error {
	return model.UpdateAboutMe(username, text)
}
