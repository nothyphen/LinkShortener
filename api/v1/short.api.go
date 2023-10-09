package v1

import (
	"linkshortner/serilizers"
	"linkshortner/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ShortAPI interface {
	Short(ctx *gin.Context)
}

type shortAPI struct {
	shortService services.ShortService
}

func NewShortAPI(ShortService services.ShortService) ShortAPI {
	return &shortAPI{
		shortService: ShortService,
	}
}

func (c *shortAPI) Short(ctx *gin.Context) {
	var request serilizers.RedirectRequest
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":"please insert shortkey",
		})
	}

	link, err := c.shortService.ShowLink(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":"not found link",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"link":link,
	})
}