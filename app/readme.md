# 业务代码开发目录


使用cwgo根据idl生成RPC代码
新建业务代码目录，
例如/app/auth，终端进入后执行：
```bash
cwgo server -I ../../idl/ --type RPC --module douyinmall --service douyinmall --idl ../../idl/auth.proto

cwgo server --type RPC --module github.com/All-Done-Right/douyin-mall-microservice/app/auth --service auth -I ../../idl --idl ../../idl/auth.proto --pass "-use github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen"
```


