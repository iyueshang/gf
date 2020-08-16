package service

import (
	"errors"
	"fmt"

	user2 "github.com/gogf/gf-demos/app/model/user"
)

func GetOne(username string) {
	user := user2.User{}
	user, err := user.GetOne(username)
	if err != nil {
		errors.New(err)
	}
	fmt.Println(user)
}
