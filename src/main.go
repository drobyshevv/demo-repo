package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Создаем роутер Gin
	r := gin.Default()

	// Обработчик для корневого пути "/"
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Привет, мир!")
	})

	// Обработчик для пути "/hello"
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	// Обработчик для JSON ответа
	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Привет в формате JSON!",
			"status":  "success",
		})
	})

	// Запускаем сервер на порту 8080
	fmt.Println("Сервер запущен на http://localhost:8080")
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
