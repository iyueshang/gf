package middleware

import (
	"github.com/gogf/gf-demos/library/response"
	"github.com/gogf/gf-jwt/example/auth"
	"github.com/gogf/gf/net/ghttp"
)

// 鉴权中间件，只有登录成功之后才能通过
func Auth(r *ghttp.Request) {
	auth.GfJWTMiddleware.MiddlewareFunc()(r)
	token, err := auth.GfJWTMiddleware.ParseToken(r)
	if err != nil {
		response.JsonExit(r, 401, "令牌失效，请重新登录", false)
	}
	r.SetCtxVar("auth_user", token.Claims)
	r.Middleware.Next()
}
