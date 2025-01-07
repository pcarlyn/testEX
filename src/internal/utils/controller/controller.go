package controller

import (
	"database/sql"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"testex/src/internal/models"
	"testex/src/internal/utils"
	"testex/src/internal/utils/database"
	"testex/src/internal/utils/database/dbconnector"
)

func GetUser(idstr string) (*models.ResponseUser, int) {

	config := database.NewConfig()
	d, err := database.ConnectToDB(config)
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		return &models.ResponseUser{}, http.StatusInternalServerError
	}
	db := dbconnector.NewDBConnector(d)

	defer d.Close()
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return &models.ResponseUser{}, http.StatusBadRequest
	}
	if id <= 0 || id > math.MaxInt {
		return &models.ResponseUser{}, http.StatusBadRequest
	}

	user, err := db.GetUserById(id)
	if err != nil {
		fmt.Printf("Error getting user from database: %v\n", err)
		return &models.ResponseUser{}, http.StatusInternalServerError
	}
	if err == sql.ErrNoRows {
		return &models.ResponseUser{}, http.StatusNotFound
	}

	responseUser := utils.UserToResponse(*user)

	return responseUser, http.StatusOK
}

func GetTopUsers() ([]models.ResponseUser, int) {
	config := database.NewConfig()
	d, err := database.ConnectToDB(config)
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		return []models.ResponseUser{}, http.StatusInternalServerError
	}
	db := dbconnector.NewDBConnector(d)

	defer d.Close()

	users, err := db.GetTopUsers()
	if err != nil {
		fmt.Printf("Error getting top users from database: %v\n", err)
		return []models.ResponseUser{}, http.StatusInternalServerError
	}

	if err == sql.ErrNoRows {
		return []models.ResponseUser{}, http.StatusNotFound
	}

	responseUsers := make([]models.ResponseUser, 0, len(users))
	for _, user := range users {
		respUser := utils.UserToResponse(user)
		responseUsers = append(responseUsers, *respUser)
	}

	return responseUsers, http.StatusOK
}

func PostCompleteTask(idstr, task string) (*models.ResponseUser, int) {

	config := database.NewConfig()
	d, err := database.ConnectToDB(config)
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		return &models.ResponseUser{}, http.StatusInternalServerError
	}
	db := dbconnector.NewDBConnector(d)

	defer d.Close()
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return &models.ResponseUser{}, http.StatusBadRequest
	}
	if id <= 0 || id > math.MaxInt {
		return &models.ResponseUser{}, http.StatusBadRequest
	}

	taskModel, err := db.GetTaskByTitle(task)
	if err != nil {
		fmt.Printf("Error getting task from database: %v\n", err)
		return &models.ResponseUser{}, http.StatusInternalServerError
	}
	if err == sql.ErrNoRows {
		return &models.ResponseUser{}, http.StatusNotFound
	}
	taskActivity, err := db.GetTaskActivityById(id)

	if err != nil {
		fmt.Printf("Error getting task activity from database: %v\n", err)
		return &models.ResponseUser{}, http.StatusInternalServerError
	}
	if err == sql.ErrNoRows {
		return &models.ResponseUser{}, http.StatusNotFound
	}

	for _, activity := range *taskActivity {
		if activity.TaskID == taskModel.ID {
			return &models.ResponseUser{}, http.StatusConflict
		}
	}

	_, err = db.PostTaskActivity(id, taskModel.ID)
	if err != nil {
		fmt.Printf("Error posting task activity to database: %v\n", err)
		return &models.ResponseUser{}, http.StatusInternalServerError
	}

	user, err := db.UpdateBalance(id, taskModel.Points)
	if err != nil {
		fmt.Printf("Error updating user balance: %v\n", err)
		return &models.ResponseUser{}, http.StatusInternalServerError
	}

	responseUser := utils.UserToResponse(*user)

	return responseUser, http.StatusOK
}

func PostReferrerCode(idstr, code string) (*models.ResponseUser, int) {

	config := database.NewConfig()
	d, err := database.ConnectToDB(config)
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		return &models.ResponseUser{}, http.StatusInternalServerError
	}
	db := dbconnector.NewDBConnector(d)

	defer d.Close()
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return &models.ResponseUser{}, http.StatusBadRequest
	}
	if id <= 0 || id > math.MaxInt {
		return &models.ResponseUser{}, http.StatusBadRequest
	}

	referrer, err := db.GerReferrerByCode(code)
	if err == sql.ErrNoRows {
		referrer, err = db.PostReferrerCode(code)
		if err != nil {
			fmt.Printf("Error posting referrer to database: %v\n", err)
			return &models.ResponseUser{}, http.StatusInternalServerError
		}
	}

	user, err := db.UpdateReferrerId(id, referrer.ID)
	if err != nil {
		fmt.Printf("Error updating user referrer: %v\n", err)
		return &models.ResponseUser{}, http.StatusInternalServerError
	}

	responseUser := utils.UserToResponse(*user)

	return responseUser, http.StatusOK
}
