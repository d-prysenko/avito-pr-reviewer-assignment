package response

import "revass/internal/model"

type UserResponse struct {
	User *model.User `json:"user"`
}
