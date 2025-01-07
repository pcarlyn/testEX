package main

import (
	"fmt"
	"testex/src/cmd/routes"
	"testex/src/internal/utils/database"

	_ "testex/docs"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func init() {
	err := database.MigrateDB(database.NewConfig())
	if err != nil {
		fmt.Printf("Error migrating database: %v\n", err)
	}
}

// @title           TestEx
// @version         1.0
// @description     API Server for User Application
// @BasePath  /
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.EchoWrapHandler(func(c *echoSwagger.Config) {
		c.URLs = []string{fmt.Sprintf("http://%s:%s/swagger/doc.json", "127.0.0.1", "8080")}
	}))

	users := e.Group("/users")

	routes.UserRoutes(users)

	e.Debug = true

	e.Logger.Fatal(e.Start(":8080"))

}
