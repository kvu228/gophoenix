package main

import (
	"pheonix/engine"
	"pheonix/example/controller/user"
)

func main() {
	phx_server := engine.NewPhoenixServer()
	api_group := phx_server.Group("/api")

	user.InitUserGroupController(api_group, "/user")
	phx_server.Run("localhost:3000")
}
