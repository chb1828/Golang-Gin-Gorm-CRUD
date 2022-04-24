package controller

import (
	"github.com/gin-gonic/gin"
	"login/service"
	"login/service/dtos"
	"net/http"
)

type LoginController interface {
	Login(ctx *gin.Context)
}

type InLoginController struct {
	loginService service.LoginService
}

func NewLoginController(service service.LoginService) LoginController {
	return &InLoginController{
		service,
	}
}

func (c *InLoginController) Login(ctx *gin.Context) {
	var parameterDTO dtos.LoginDTO

	// 들어온 json 형식을 검증한다
	if err := ctx.ShouldBindJSON(&parameterDTO); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, "올바르지 않은 형태입니다")
		return
	}
	if !c.loginService.Check(parameterDTO) {
		ctx.JSON(http.StatusUnauthorized, "알 수없는 사용자 입니다")
		return
	}

	token, err := c.loginService.CreateToken(parameterDTO)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, token)
}
