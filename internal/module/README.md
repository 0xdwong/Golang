# go module 
go module 是Go语言从 1.11 版本之后官方推出的版本管理工具，并且从 Go1.13 版本开始，go module 成为了Go语言默认的依赖管理工具。


## 使用:
1.  golang 升级到 1.11 版本以上
2.  设置 GO111MODULE

    export GO111MODULE=on

在项目中使用

1. go mod init <packageName>
2. 添加依赖
   
   比如：
    ```go get go.uber.org/zap```

3. 执行 go run main.go 运行代码, 会自动查找依赖自动下载



## 常用的go mod命令

| 命令             | 作用 |
| :---            | :---|
| go mod download |	下载依赖包到本地（默认为 GOPATH/pkg/mod 目录）|
| go mod edit	  |  编辑 go.mod 文件|
| go mod graph	  |  打印模块依赖图|
| go mod init	  |  初始化当前文件夹，创建 go.mod 文件|
| go mod tidy	  |  增加缺少的包，删除无用的包|
| go mod vendor	  |  将依赖复制到 vendor 目录下|
| go mod verify	  |  校验依赖|
| go mod why	  |  解释为什么需要依赖|

