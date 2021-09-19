package controller

import (
	"github.com/gin-gonic/gin"
	"login/service"
	"login/service/dtos"
	"net/http"
)

type UserController interface {
	AddUser(*gin.Context)
	RemoveUser(ctx *gin.Context)
	Select(ctx *gin.Context)
	SelectAll(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userController{
		service,
	}
}

func (u userController) AddUser(ctx *gin.Context) {
	var user dtos.UserDTO

	// 들어온 json 형식을 검증한다
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, "올바르지 않은 형태입니다")
		return
	}

	resultUser, err := u.userService.Save(user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error":"User를 저장하던 도중 에러가 발생했습니다."})
		return
	}

	ctx.JSON(http.StatusCreated,gin.H{"data":resultUser})

}

func (u userController) RemoveUser(ctx *gin.Context) {
	id := ctx.PostForm("id")
	if id =="" {
		return
	}
	err := u.userService.Delete(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User를 삭제하던 도중 에러가 발생했습니다."})
	}

	ctx.JSON(http.StatusOK,gin.H{})
}

func (u userController) Select(ctx *gin.Context) {
	id := ctx.Param("id")
	if id=="" {
		ctx.JSON(http.StatusBadRequest,gin.H{"error":"잘못된 요청입니다"})
		return
	}
	user, err := u.userService.Select(id)
	if err !=nil {
		ctx.JSON(http.StatusNotFound,gin.H{"error":"사용자를 찾을 수 없습니다"})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"data":user})
}

func (u userController) SelectAll(ctx *gin.Context) {
	users := u.userService.SelectAll()
	ctx.JSON(http.StatusOK,gin.H{"data":users})
}