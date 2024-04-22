package main

import (
		"github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "todo-app/backend/controllers"
)

func main() {
		gin.SetMode(gin.ReleaseMode)
    router := gin.Default()

		// すべてのドメインとすべてのメソッドに対するCORS
		router.Use(cors.New(cors.Config{
			AllowOrigins: []string{"*"},  // すべてのオリジンからのアクセスを許可
			AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
			ExposeHeaders: []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge: 12 * 60 * 60,  // 12 hours
		}))

		// ルートパスのハンドラーを追加
		router.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
					"message": "Welcome to the ToDo API!",
			})
		})

    router.GET("/todos", controllers.GetTodos)
    router.POST("/todos", controllers.AddTodo)
    router.DELETE("/todos/:id", controllers.DeleteTodo)

    router.Run(":8080")
}
