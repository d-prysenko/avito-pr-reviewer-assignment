package request

type SetIsActiveRequest struct {
	UserID   string `json:"user_id" validate:"required,max=32"`
	IsActive bool   `json:"is_active" validate:"boolean"`
}
