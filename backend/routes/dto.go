package routes

type Response struct {
	Code   int         `json:"code"`
	ErrMsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}
