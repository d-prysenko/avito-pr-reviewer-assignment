package model

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Team     int64  `json:"team"`
	IsActive bool   `json:"is_active"`
}

type UserPullRequest struct {
	PullRequestID   string `json:"pull_request_id"`
	PullRequestName string `json:"pull_request_name"`
	AuthorID        string `json:"author_id"`
	Status          string `json:"status"`
}
