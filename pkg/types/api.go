package types

//type ApiResult struct {
//	Code    int      `json:"code"`
//	Message string   `json:"message"`
//	Data    LoginRes `json:"data"`
//}

type ApiResultInter struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	//Data    LoginRes `json:"data"`
	Data interface{} `json:"data"`
}

type ApiResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	//Data    LoginRes `json:"data"`
	Data []byte `json:"data"`
}

type LoginRes struct {
	Token  string `json:"token"`
	UserID string `json:"user_id"`
}
