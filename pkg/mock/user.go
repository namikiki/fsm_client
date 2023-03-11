package mock

import (
	"time"

	"fsm_client/pkg/ent"
	"fsm_client/pkg/types"
)

func NewUser() ent.User {
	return ent.User{
		ParentID:   "",
		UserID:     "",
		Name:       "",
		Deleted:    false,
		CreateTime: time.Time{},
		ModTime:    time.Time{},
	}
}

func NewAccount() types.UserLoginReq {
	return types.UserLoginReq{
		Email:    "1231231231@gmail.com",
		PassWord: "1231der232342423",
	}
}

func NewRegis() types.UserRegister {
	return types.UserRegister{
		Email:    "1231231231@gmail.com",
		PassWord: "1231der232342423",
		UserName: "zylzyl",
	}
}
