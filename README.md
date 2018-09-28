## 交叉编译
* linux 64位
``` bash
gox -osarch="linux/amd64" -output="./tmp/webhook"

```
## 开发测试
``` bash
go run main.go -config-file=./config.yml -static-dir=./static/

```