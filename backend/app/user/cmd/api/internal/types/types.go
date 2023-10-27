// Code generated by goctl. DO NOT EDIT.
package types

type RegisterReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}

type RegisterResp struct {
	Status  int    `json:"status"`
	Data    string `json:"data"`
	Message string `json:"msg"`
	Error   string `json:"error"`
}

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
}

type Data struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResp struct {
	Status  int    `json:"status"`
	Data    Data   `json:"data"`
	Message string `json:"msg"`
	Error   string `json:"error"`
}
