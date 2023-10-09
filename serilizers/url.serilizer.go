package serilizers

type ShortRequest struct {
	Link string `json:"link" form:"link" binding:"required,min=1"`
}