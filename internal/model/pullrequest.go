package model

type PullRequest struct {
	PullRequestID     string   `json:"pull_request_id"`
	PullRequestName   string   `json:"pull_request_name"`
	AuthorID          string   `json:"author_id"`
	Status            string   `json:"status"`
	MergedAt          *string   `json:"merged_at,omitempty"`
	AssignedReviewers []string `json:"assigned_reviewers"`
}
