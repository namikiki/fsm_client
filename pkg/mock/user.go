package mock

import (
	"fsm_client/pkg/types"
)

func NewLogin() types.UserLoginReq {
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
