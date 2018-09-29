## 交叉编译
* linux 
具体看 Makefile
``` bash
make linux

```
## 开发测试
``` bash
go run ./api/main.go -config-file=./config.simple.yml -static-dir=./static/

```
## gogs 添加 Web 钩子

推送地址: http://ip:9090/hook/jobname

数据格式: application/json

密钥文本: 与config.simple.yml 内对应
