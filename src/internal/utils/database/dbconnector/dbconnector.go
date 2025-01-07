package dbconnector

import (
	"database/sql"
	"testex/src/internal/models"
)

type DataBase struct {
	DB *sql.DB
}

func NewDBConnector(db *sql.DB) *DataBase {
	return &DataBase{
		DB: db,
	}
}

func (d *DataBase) GetUserById(id int) (*models.User, error) {
	query := "SELECT id, telegram_id, active, registered_at, user_name, status_id, last_visit, balance, isadmin, referrer_id FROM users WHERE id = $1"

	row := d.DB.QueryRow(query, id)

	user := &models.User{}

	err := row.Scan(&user.ID, &user.TelegramID, &user.Active, &user.RegisteredAt, &user.UserName, &user.StatusID, &user.LastVisit, &user.Balance, &user.IsAdmin, &user.ReferrerID)
	if err != nil {
		return &models.User{}, err
	}
	return user, nil
}

func (d *DataBase) GetTopUsers() ([]models.User, error) {
	query := "SELECT id, telegram_id, active, registered_at, user_name, status_id, last_visit, balance, isadmin, referrer_id FROM users ORDER BY balance DESC LIMIT $1"

	rows, err := d.DB.Query(query, 3)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]models.User, 0)
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.ID, &user.TelegramID, &user.Active, &user.RegisteredAt, &user.UserName, &user.StatusID, &user.LastVisit, &user.Balance, &user.IsAdmin, &user.ReferrerID)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (d *DataBase) GetTaskByTitle(title string) (*models.Task, error) {
	query := "SELECT id, title, points FROM tasks WHERE title = $1"

	row := d.DB.QueryRow(query, title)

	task := &models.Task{}

	err := row.Scan(&task.ID, &task.Title, &task.Points)
	if err != nil {
		return &models.Task{}, err
	}
	return task, nil
}

func (d *DataBase) GetTaskActivityById(userId int) (*[]models.TaskActivity, error) {
	query := "SELECT id, user_id, task_id FROM tasks_activity WHERE user_id = $1"

	rows, err := d.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	taskActivities := make([]models.TaskActivity, 0)
	for rows.Next() {
		taskActivity := models.TaskActivity{}
		err := rows.Scan(&taskActivity.ID, &taskActivity.UserID, &taskActivity.TaskID)
		if err != nil {
			return nil, err
		}
		taskActivities = append(taskActivities, taskActivity)
	}
	return &taskActivities, nil
}

func (d *DataBase) PostTaskActivity(userId, taskId int) (int, error) {
	query := "INSERT INTO tasks_activity (user_id, task_id) VALUES ($1, $2) RETURNING id"

	row := d.DB.QueryRow(query, userId, taskId)

	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (d *DataBase) UpdateBalance(userId, balance int) (*models.User, error) {
	query := "UPDATE users SET balance = balance + $1 WHERE id = $2 RETURNING id, telegram_id, active, registered_at, user_name, status_id, last_visit, balance, isadmin, referrer_id"

	row := d.DB.QueryRow(query, balance, userId)

	user := &models.User{}

	err := row.Scan(&user.ID, &user.TelegramID, &user.Active, &user.RegisteredAt, &user.UserName, &user.StatusID, &user.LastVisit, &user.Balance, &user.IsAdmin, &user.ReferrerID)
	if err != nil {
		return &models.User{}, err
	}
	return user, nil
}

func (d *DataBase) GerReferrerByCode(code string) (*models.Referrer, error) {
	query := "SELECT id, code FROM referrers WHERE code = $1"

	row := d.DB.QueryRow(query, code)

	referrer := &models.Referrer{}

	err := row.Scan(&referrer.ID, &referrer.Code)
	if err != nil {
		return &models.Referrer{}, err
	}
	return referrer, nil
}

func (d *DataBase) PostReferrerCode(code string) (*models.Referrer, error) {
	query := "INSERT INTO referrers (code) VALUES ($1) RETURNING id, code"

	row := d.DB.QueryRow(query, code)

	referrer := &models.Referrer{}

	err := row.Scan(&referrer.ID, &referrer.Code)
	if err != nil {
		return &models.Referrer{}, err
	}
	return referrer, nil
}

func (d *DataBase) UpdateReferrerId(userId, referrerId int) (*models.User, error) {
	query := "UPDATE users SET referrer_id = $1 WHERE id = $2 RETURNING id, telegram_id, active, registered_at, user_name, status_id, last_visit, balance, isadmin, referrer_id"

	row := d.DB.QueryRow(query, referrerId, userId)

	user := &models.User{}

	err := row.Scan(&user.ID, &user.TelegramID, &user.Active, &user.RegisteredAt, &user.UserName, &user.StatusID, &user.LastVisit, &user.Balance, &user.IsAdmin, &user.ReferrerID)
	if err != nil {
		return &models.User{}, err
	}
	return user, nil
}
