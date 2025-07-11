package converter

import (
	"github.com/fernanda-syafalam/backend-monitoring-notification/internal/entity"
	"github.com/fernanda-syafalam/backend-monitoring-notification/internal/model"
)

func UserToResponse(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		ID: user.ID,
		Name: user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func UserToTokenResponse(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		Token: user.Token,
	}
}


func UserToEvent(user *entity.User) *model.UserEvent {
	return &model.UserEvent{
		ID: user.ID,
		Name: user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}