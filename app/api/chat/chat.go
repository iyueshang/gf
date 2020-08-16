package chat

import "github.com/gogf/gf/net/ghttp"

func Index(r *ghttp.Request) {
	_ = r.Response.WriteTpl("chat/index.html")
}
