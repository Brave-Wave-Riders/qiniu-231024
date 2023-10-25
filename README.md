# qiniu-231024

```shell
goctl api -o video.api

api生成
goctl api go -api video.api -dir .

生成proto文件模板
goctl rpc template -o video.proto

rpc生成
goctl rpc protoc video.proto --go_out=. --go-grpc_out=. --zrpc_out=.

```


![img.png](img.png)