package main

import (
    "github.com/gin-gonic/gin"
    "todo-app/backend/controllers"
)

func main() {
		gin.SetMode(gin.ReleaseMode)
    router := gin.Default()

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
