package types

type UserRegister struct {
	Email    string `json:"email,omitempty" validate:"required,min=10,email"`
	PassWord string `json:"password,omitempty" validate:"required,min=10"`
	UserName string `json:"username,omitempty" validate:"required,min=5"`
}

type UserLoginReq struct {
	Email    string `json:"email" validate:"required"`
	PassWord string `json:"password" validate:"required"`
}

type UserLoginRes struct {
	Token  string `json:"token"`
	UserID string `json:"user_id"`
}
