package main

import (
	"fmt"
	"net/http"
	"pheonix/engine"
)

func MiddlewareAuthorization() engine.HandlerFunc {
	handler := func(ctx *engine.PhoenixContext) {
		fmt.Println("middleware authorization running")
		if ctx.GetParam("id") == "42" {
			ctx.Next()
			// backtracking
		} else {
			ctx.Halt()
			ctx.JSON(http.StatusUnauthorized, "Not Authorized this user")
			return
		}
	}
	return handler
}

func main() {
	// cấp phát 1 server mới
	app := engine.NewPhoenixServer()
	//api_group := app.Group("/api")
	//api_group.Use()SecurityControllerHandler()
	//group_api_user := api_group.Group("/user")
	//group_api_product := api_group.Group("/product")

	// nhóm các api có tiền tố "/api"
	api_group := app.Group("/api")
	//các api trong api_group phải dùng middlewareAuthorization trước khi tới API endpoint
	api_group.Use(MiddlewareAuthorization())

	api_group.GET("/v1/user/:id", func(ctx *engine.PhoenixContext) {
		ctx.JSON(200, map[string]any{
			"user_id": ctx.GetParam("id"),
		})
	})

	// Server chạy trên localhost:8000
	app.Run("localhost:8000")
}
