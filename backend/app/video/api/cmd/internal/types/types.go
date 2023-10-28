// Code generated by goctl. DO NOT EDIT.
package types

type UploapReq struct {
	Title string `form:"title"`
	Kind  uint   `form:"kind"`
}

type UploapResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type DeleteReq struct {
	Id int32 `json:"id"`
}

type DeleteResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type StarReq struct {
	Id int32 `json:"id"`
}

type Data struct {
	Star int64 `json:"star"`
	Id   int64 `json:"id"`
}

type StarResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data Data   `json:"data"`
}