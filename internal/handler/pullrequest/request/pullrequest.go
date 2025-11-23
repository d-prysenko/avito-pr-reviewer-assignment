package request

type PullRequestCreateRequest struct {
	PullRequestID   string `json:"pull_request_id" validate:"required,max=32"`
	PullRequestName string `json:"pull_request_name" validate:"required,max=128"`
	AuthorID        string `json:"author_id" validate:"required,max=32"`
}

type PullRequestMergeRequest struct {
	PullRequestID string `json:"pull_request_id" validate:"required,max=32"`
}

type PullRequestReassignRequest struct {
	PullRequestID string `json:"pull_request_id" validate:"required,max=32"`
	OldReviewerID  string `json:"old_reviewer_id" validate:"required,max=32"`
}
