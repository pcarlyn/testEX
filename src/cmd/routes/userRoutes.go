package routes

import (
	"testex/src/cmd/handlers"

	"github.com/labstack/echo/v4"
)

func UserRoutes(group *echo.Group) {
	group.GET("/:id/status", handlers.GetStatus)
	group.GET("/leaderboard", handlers.GetLeaderboard)
	group.POST("/:id/task/complete", handlers.PostCompleteTask)
	group.POST("/:id/referrer", handlers.PostReferrerCode)
}
