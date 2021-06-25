package controller

import (
	"mock-server/service"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	// app
	appController *AppController

	// auth app tag
	// authAppTagController *AuthAppTagController
}

func (c *Controller) ApplyRoutes(e *gin.Engine) {
	c.ApplyAuthAppRoutes(e)

	c.ApplyAuthAppTagRoutes(e)

	// v1 := e.Group("/v1", h.AuthRequiredMiddleware())
	// v1.GET("/application", h.ListApplication)
	// v1.GET("/application/:app", h.GetApplication)
	// v1.GET("/appinstance", h.ListAppInstance)
	// v1.GET("/appinstance/:ins", h.ListAppInstance)
}

func (c *Controller) ApplyAuthAppRoutes(e *gin.Engine) {
	// auth app
	gm_api_v1 := e.Group("/gm/api/v1")
	gm_api_v1.GET("/apps", c.appController.getApps)
	gm_api_v1.GET("/app/:app_id/gms", c.appController.getGMs)
	gm_api_v1.GET("/app/:app_id/groups", c.appController.getGroups)
	gm_api_v1.POST("/app/:app_id/gm", c.appController.addGMToGroup)
	gm_api_v1.GET("/gm_detail", c.appController.getUserInfo)
	gm_api_v1.PATCH("/gm_detail", c.appController.updateUserInfo)
}

func (c *Controller) ApplyAuthAppTagRoutes(e *gin.Engine) {
	// api := e.Group("/gm/api/v1")
}

func NewController() *Controller {
	return &Controller{
		&AppController{
			service.NewAppJSONService(),
		},
		// &AuthAppTagController{
		// 	service.NewAuthAppTagJSONService(),
		// },
	}
}
