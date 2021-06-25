package controller

import (
	"fmt"
	"mock-server/model"
	"mock-server/service"

	"github.com/gin-gonic/gin"
)

type AppController struct {
	appService service.AppServiceInterface
}

func (a *AppController) getApps(ctx *gin.Context) {
	apps, err := a.appService.GetApps(ctx)
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	fmt.Println(apps)
	ctx.JSON(200, "ok")
}

func (a *AppController) getGMs(ctx *gin.Context) {
	fmt.Println("[query params / page]", ctx.Query("page"))

	gms, err := a.appService.GetGMs(ctx)
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.Data(200, "application/json", []byte(gms))
}

func (a *AppController) getGroups(ctx *gin.Context) {
	groups, err := a.appService.GetGroups(ctx)
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.Data(200, "application/json", []byte(groups))
}

func (a *AppController) addGMToGroup(ctx *gin.Context) {
	var req model.AddUserToGroupReq
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(500, err)
	}

	fmt.Println("[addGMToGroup req] ", req)

	res := `{
		"error": {},
	}`
	ctx.Data(200, "application/json", []byte(res))
}

func (a *AppController) getUserInfo(ctx *gin.Context) {
	fmt.Println("GET params gm_id", ctx.Query("gm_id"))

	userInfo, err := a.appService.GetUserInfo(ctx)
	if err != nil {
		ctx.JSON(500, err)
		fmt.Println(err)
		return
	}

	ctx.Data(200, "application/json", []byte(userInfo))
}

func (a *AppController) updateUserInfo(ctx *gin.Context) {
	fmt.Println("haha")

	var userInfo model.UserInfo
	err := ctx.BindJSON(&userInfo)
	if err != nil {
		ctx.JSON(500, err)
	}

	fmt.Println("updateUserInfo post data", userInfo)

	res := `{
		"error": {}
	}`
	ctx.Data(200, "application/json", []byte(res))
}
