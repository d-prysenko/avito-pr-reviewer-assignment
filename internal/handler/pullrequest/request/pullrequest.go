package request

type PullRequestCreateRequest struct {
	PullRequestID   string `json:"pull_request_id" validate:"required,max=128"`
	PullRequestName string `json:"pull_request_name" validate:"required,max=128"`
	AuthorID        string `json:"author_id" validate:"required,max=128"`
}

type PullRequestMergeRequest struct {
	PullRequestID   string `json:"pull_request_id" validate:"required,max=128"`
}