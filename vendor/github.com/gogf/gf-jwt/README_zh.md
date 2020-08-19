# gf-jwt
gf 的 jwt 插件。

这个插件是 fork 了 [https://github.com/appleboy/gin-jwt](https://github.com/appleboy/gin-jwt) 插件,修改为 [https://github.com/gogf/gf](https://github.com/gogf/gf) 插件.


[英文](README.md) [中文](README_zh.md)


## 使用

下载安装

```sh
$ go get github.com/gogf/gf-jwt
```

导入

```go
import "github.com/gogf/gf-jwt"
```

## 例子

查看示例 [example/auth/auth.go](example/auth/auth.go) , 使用 `ExtractClaims` 可以自定义用户数据.

[embedmd]:# (example/auth/auth.go go)

## Demo

运行 `example/server/server.go` 在 `8000`端口.

```bash
$ go run example/server/server.go
```

![api screenshot](screenshot/server.png)

通过 [httpie](https://github.com/jkbrzt/httpie) ,在命令行来测试下效果.

### 登录接口:

```bash
$ http -v --form  POST localhost:8000/login username=admin password=admin
```

命令行输出

![api screenshot](screenshot/login.png)

### 刷新 token 接口:

```bash
$ http -v -f GET localhost:8000/user/refresh_token "Authorization:Bearer xxxxxxxxx"  "Content-Type: application/json"
```

命令行输出

![api screenshot](screenshot/refresh_token.png)

### hello 接口

我们使用用户名 `admin` 和密码 `admin` 测试一下 hello 接口的返回

```bash
$ http -f GET localhost:8000/user/hello "Authorization:Bearer xxxxxxxxx"  "Content-Type: application/json"
```

命令行输出

![api screenshot](screenshot/hello.png)

### 用户验证接口

我们用未授权的 token 来测试 hello 接口的返回

```bash
$ http -f GET localhost:8000/user/hello "Authorization:Bearer xxxxxxxxx"  "Content-Type: application/json"
```

命令行输出

![api screenshot](screenshot/401.png)


再次感谢[https://github.com/appleboy/gin-jwt](https://github.com/appleboy/gin-jwt)。