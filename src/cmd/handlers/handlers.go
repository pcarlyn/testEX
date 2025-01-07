package handlers

import (
	"net/http"
	"testex/src/internal/models"
	"testex/src/internal/utils/controller"

	"github.com/labstack/echo/v4"
)

// Handler for get User infomation by id
// @Summary Summary
// @Description Description
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.ResponseUser "Success Response"
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 404 {object} models.ErrorResponse "Not Found"
// @Failure 500 {object} models.ErrorResponse "Server error"
// @Router /users/{id}/status [get]
func GetStatus(c echo.Context) error {
	id := c.Param("id")

	user, statusCode := controller.GetUser(id)
	switch statusCode {
	case http.StatusBadRequest:
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Code: "400", Error: "Bad Request"})
	case http.StatusNotFound:
		return c.JSON(http.StatusNotFound, models.ErrorResponse{Code: "404", Error: "User not found"})
	case http.StatusInternalServerError:
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{Code: "500", Error: "Internal Server Error"})
	}
	return c.JSON(statusCode, user)
}

// Handler for get Leaderboard
// @Summary Summary
// @Description Description
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} []models.ResponseUser "Success Response"
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 404 {object} models.ErrorResponse "Not Found"
// @Failure 500 {object} models.ErrorResponse "Server error"
// @Router /users/leaderboard [get]
func GetLeaderboard(c echo.Context) error {

	users, statusCode := controller.GetTopUsers()
	switch statusCode {
	case http.StatusBadRequest:
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Code: "400", Error: "Bad Request"})
	case http.StatusNotFound:
		return c.JSON(http.StatusNotFound, models.ErrorResponse{Code: "404", Error: "User not found"})
	case http.StatusInternalServerError:
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{Code: "500", Error: "Internal Server Error"})
	}
	return c.JSON(statusCode, users)
}

// Handler for Post Complete Task
// @Summary Summary
// @Description Description
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param task body models.PostResponse true "Task"
// @Success 200 {object} models.ResponseUser "Success Response"
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 404 {object} models.ErrorResponse "Not Found"
// @Failure 500 {object} models.ErrorResponse "Server error"
// @Router /users/{id}/task/complete [post]
func PostCompleteTask(c echo.Context) error {

	var resp models.PostResponse
	idstr := c.Param("id")
	err := c.Bind(&resp)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Error parsing request body")
	}
	user, statusCode := controller.PostCompleteTask(idstr, resp.Task)
	switch statusCode {
	case http.StatusBadRequest:
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Code: "400", Error: "Bad Request"})
	case http.StatusNotFound:
		return c.JSON(http.StatusNotFound, models.ErrorResponse{Code: "404", Error: "User not found"})
	case http.StatusInternalServerError:
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{Code: "500", Error: "Internal Server Error"})
	}
	return c.JSON(statusCode, user)
}

// Handler for Post Referrer Code
// @Summary Summary
// @Description Description
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param code body models.PostReferrer true "Referrer Code"
// @Success 200 {object} models.ResponseUser "Success Response"
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 404 {object} models.ErrorResponse "Not Found"
// @Failure 500 {object} models.ErrorResponse "Server error"
// @Router /users/{id}/referrer [post]
func PostReferrerCode(c echo.Context) error {
	var resp models.PostReferrer
	idstr := c.Param("id")

	err := c.Bind(&resp)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Error parsing request body")
	}
	user, statusCode := controller.PostReferrerCode(idstr, resp.Code)
	switch statusCode {
	case http.StatusBadRequest:
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Code: "400", Error: "Bad Request"})
	case http.StatusNotFound:
		return c.JSON(http.StatusNotFound, models.ErrorResponse{Code: "404", Error: "User not found"})
	case http.StatusInternalServerError:
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{Code: "500", Error: "Internal Server Error"})
	}
	return c.JSON(statusCode, user)
}
