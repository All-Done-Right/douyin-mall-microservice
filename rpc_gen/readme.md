在rpcgen里面使用cwgo生成client代码

```bash
cwgo client --type RPC --module github.com/All-Done-Right/douyin-mall-microservice/rpc_gen --service auth -I ../idl --idl ../idl/auth.proto
```