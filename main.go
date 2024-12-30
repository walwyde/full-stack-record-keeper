package main

import (
	"fmt"
	"go-api/routes"
	"os"
	"path/filepath"
	"github.com/gin-gonic/gin"
)


func main() {
	executable, err := os.Executable()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	routesFilePath := filepath.Join(filepath.Dir(executable), "routes.go")
	fmt.Println("File path:", routesFilePath)

	app := routes.SetupRouter()

	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": c.Request.UserAgent(),
		})
	})

	routes.RunApp()

	app.Run(":8080") // listen and serve on port 8080
}