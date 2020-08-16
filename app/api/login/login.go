package login

import "github.com/gogf/gf/net/ghttp"

func Index(r *ghttp.Request) {

	r.Response.WriteTpl("login/index.html")
}
