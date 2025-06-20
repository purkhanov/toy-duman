package controller

import (
	"net/http"
	"toy-duman/web/service"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
	userService service.UserService
}

func NewIndexController(g *gin.RouterGroup) *IndexController {
	a := &IndexController{}
	a.initRouter(g)

	return a
}

func (c *IndexController) initRouter(g *gin.RouterGroup) {
	g.GET("/", c.index)
}

func (c *IndexController) index(ctx *gin.Context) {
	str, _ := c.userService.GetUser()

	ctx.JSON(http.StatusOK, gin.H{"msg": str})
}
