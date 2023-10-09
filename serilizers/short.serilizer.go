package serilizers

type RedirectRequest struct {
	Short string `json:"short" form:"short" binding:"required,min=1"`
}