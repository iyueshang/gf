package user

import (
	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf-demos/app/model/user"
	"github.com/gogf/gf-demos/library/response"
	"github.com/gogf/gf/net/ghttp"
)

func Index(r *ghttp.Request) {

}

func Profile(r *ghttp.Request) {
	auth_user := r.GetCtxVar("auth_user")
	user := user.User{}
	profile, err := user.GetProfile(gconv.Int8("1"))
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "成功", profile)
}
