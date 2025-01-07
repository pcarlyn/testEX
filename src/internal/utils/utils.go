package utils

import "testex/src/internal/models"

func UserToResponse(user models.User) *models.ResponseUser {
	return &models.ResponseUser{
		ID:           user.ID,
		TelegramID:   user.TelegramID,
		Active:       user.Active,
		RegisteredAt: user.RegisteredAt,
		UserName:     user.UserName,
		StatusID:     user.StatusID,
		LastVisit:    user.LastVisit,
		Balance:      user.Balance,
		IsAdmin:      user.IsAdmin,
		ReferrerID:   user.ReferrerID,
		Cache:        false,
	}
}
