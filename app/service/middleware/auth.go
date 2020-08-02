package middleware

import (
	"github.com/gogf/gf-jwt/example/auth"
	"github.com/gogf/gf/net/ghttp"
)

// 鉴权中间件，只有登录成功之后才能通过
func Auth(r *ghttp.Request) {
	auth.GfJWTMiddleware.MiddlewareFunc()(r)
	r.Middleware.Next()
}
