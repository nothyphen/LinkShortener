package v1

import (
	"linkshortner/serilizers"
	"linkshortner/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LinkApi interface {
	Link(ctx *gin.Context)
}

type linkAPI struct {
	linkServices services.LinkService
}

func NewLinkApi(lkservice services.LinkService) LinkApi {
	return &linkAPI{
		linkServices: lkservice,
	}
}

func (c *linkAPI) Link(ctx *gin.Context) {
	var request serilizers.ShortRequest
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":"error",
		})
	}

	shortkey, err := c.linkServices.AddLinkService(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":"error",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"key":shortkey,
	})
}