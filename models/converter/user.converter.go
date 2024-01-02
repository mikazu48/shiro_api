package converter

import "shiro_api/models"

func UserToResponse(user *models.User) *models.UserResponse {
	return &models.UserResponse{
		Username: user.Username,
		Email:    user.Email,
		FullName: user.FullName,
	}
}
