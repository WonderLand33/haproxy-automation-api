# haproxy-automation-api

haproxy-automation-api 是一个使用 Golang 编写的基于 HTTP 的 API 服务，能够帮助 HAProxy 用户自动更新 HAProxy 的 IP 黑名单列表，并且自动重启(热更新) HAProxy。


## 依赖

``` golang
	github.com/labstack/echo/v4 v4.10.2
	github.com/scylladb/go-set v1.0.2
	github.com/spf13/cobra v1.7.0
```

## 功能 

- [x] 封禁IP
- [x] 移除某个封禁的IP
- [x] 列出所有封禁IP


## 安装和使用

要安装和使用 haproxy-automation-api，请执行以下步骤：


1. 克隆或下载本代码库；

2. 进入项目根目录，执行以下命令，下载所需的依赖项：

```go
go mod tidy -v
```

3. 将项目编译为可执行文件
``` sh
make build
```

4. 将可执行文件发送到Linux服务器
```sh
rsync -zaP ./dist/haproxy-automation-api-linux-x86 username@address:/etc/systemd/system/
```

5. 运行
``` sh
# 运行
systemctl start haproxy-automation-api
# 重启
systemctl restart haproxy-automation-api
# 状态
systemctl status haproxy-automation-api
```


**本地运行：**

``` go
go run main.go
```

在默认情况下， `23333` 端口上运行。如果您希望将其部署到生产环境中，请使用适当的方式更改端口。


## API 文档


- **POST /api/banned-ip**

添加或删除 IP 地址到 HAProxy 的黑名单列表中，同时支持对 HAProxy 进行重启。


例如，若要添加 IP 地址，您可以执行以下命令：

``` bash
curl -X POST -H "Content-Type: application/json" -d '{"IP": "192.168.1.2"}' http://localhost:8080/api/banned_ips
```


- **DELETE /api/banned-ip**

删除 HAProxy 的黑名单列表中的 IP 地址，同时支持对 HAProxy 进行重启。请求体需要包含以下选项：

例如，若要删除一个 IP 地址，您可以执行以下命令：

``` bash
curl -X DELETE -H "Content-Type: application/json" -d '{"IP": "192.168.1.2"}' http://localhost:8080/api/banned-ip
```

- **GET /api/banned-ips**

查看 HAProxy 的黑名单列表


## 注意事项


• 在开始使用 haproxy-automation-api 之前，请确保您完全了解 API 的工作原理，并在开发环境中测试其所有功能。

• 禁止在生产环境中使用本项目之前请先测试，以免数据丢失或其他异常。