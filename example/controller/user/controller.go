package user

import (
	"fmt"
	"pheonix/engine"
)

func InitUserGroupController(parentGroup *engine.Engine, apiPath string) {
	user_group := parentGroup.Group(apiPath)
	LoginUser(user_group)
}

func LoginUser(user_group *engine.Engine) {
	type UserHeaderReq struct {
		AuthToken string `header:"auth_token,required"`
		UserId    string `header:"user_id"`
		UserName  string `header:"name"`
	}

	user_group.GET("/register", func(ctx *engine.PhoenixContext) {
		var userHeader UserHeaderReq
		ctx.ParseRequestHeader(&userHeader)
		fmt.Println(userHeader.AuthToken)
		fmt.Println(userHeader.UserId)
		fmt.Println(userHeader.UserName)
	})
}
