package response

import "revass/internal/model"

type UserResponse struct {
	User *model.User `json:"user"`
}

type UserReviewResponse struct {
	UserID       string                   `json:"user_id"`
	PullRequests []*model.UserPullRequest `json:"pull_requests"`
}
