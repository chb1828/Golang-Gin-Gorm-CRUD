package route

import (
	"github.com/gin-gonic/gin"
	"login/config/middleware"
	"login/controller"
	"login/core/database"
	"login/core/repository"
	"login/service"
)

var (
	loginController controller.LoginController
	userController controller.UserController

)

func InitRoutes(server *gin.Engine) {

	// 미들웨어를 적용한다
	server.Use(gin.Recovery(), middleware.CORSMiddleware(), middleware.Logger())

	// Controller를 초기화 한다
	initController()

	apiRoutes := server.Group("/api")
	{
		apiRoutes.POST("/login",func(ctx *gin.Context) {
			loginController.Login(ctx)
		})

		apiRoutes.POST("/user",func(ctx *gin.Context) {
			userController.AddUser(ctx)
		})

		apiRoutes.DELETE("/user",func(ctx *gin.Context) {
			userController.RemoveUser(ctx)
		})

		apiRoutes.GET("/user/:id", func(ctx *gin.Context) {
			userController.Select(ctx)
		})

	}
}

func initController() {
	db := database.GetDB()
	loginService := service.NewLoginService()
	loginController = controller.NewLoginController(loginService)
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController = controller.NewUserController(userService)
}
