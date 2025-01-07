package main

import (
	"fmt"
	"testex/src/cmd/handlers"
	"testex/src/cmd/routes"
	"testex/src/internal/models"
	"testex/src/internal/utils/database"

	_ "testex/docs"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/swagger/*", echoSwagger.EchoWrapHandler(func(c *echoSwagger.Config) {
		c.URLs = []string{fmt.Sprintf("http://%s:%s/swagger/doc.json", "127.0.0.1", "8080")}
	}))

	e.POST("/login", handlers.PostUserGetJWT)

	users := e.Group("/users")

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(models.JwtCustomClaims)
		},
		SigningKey: []byte("TestEx"),
	}

	users.Use(echojwt.WithConfig(config))

	routes.UserRoutes(users)

	e.Debug = true

	e.Logger.Fatal(e.Start(":8080"))

}
