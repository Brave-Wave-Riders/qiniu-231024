### 1. "upload_vedio"

1. route definition

- Url: /api/v1/vedio/upload_vedio
- Method: POST
- Request: `UploapReq`
- Response: `UploapResp`

2. request definition



```golang
type UploapReq struct {
	Title string `form:"title"`
	Kind uint `form:"kind"`
}
```


3. response definition



```golang
type UploapResp struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}
```

