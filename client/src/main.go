package main

import (
	"client/src/controllers"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	// 前端網站
	r.LoadHTMLFiles(
		"./src/views/index.html",
	)
	// 當有人連進來時會呼叫 controller 裡面的 SendHTML
	r.GET("/room", controllers.SendHTML)
	// 只要有人連進來就會呼叫 controller 裡面的 ConnectWebSocket
	r.GET("/room/ws", controllers.ConnectWebSocket)

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8000
	r.Run(":8000")
}
