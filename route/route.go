package route

import (
	"github.com/gin-gonic/gin"
	"login/config/middleware"
	"login/controller"
	"login/core/database"
	"login/core/repository"
	"login/module"
	"login/service"
)

var (
	loginController controller.LoginController
	userController  controller.UserController
)

func InitRoutes(server *gin.Engine) {

	// 미들웨어를 적용한다
	server.Use(gin.Recovery(), middleware.CORSMiddleware(), middleware.Logger())

	// Controller를 초기화 한다
	initController()

	apiRoutes := server.Group("/api")
	{
		apiRoutes.POST("/login", func(ctx *gin.Context) {
			loginController.Login(ctx)
		})

		userRoutes := apiRoutes.Group("/user")
		{
			userRoutes.POST("/", func(ctx *gin.Context) {
				userController.AddUser(ctx)
			})

			userRoutes.DELETE("/", func(ctx *gin.Context) {
				userController.RemoveUser(ctx)
			})

			userRoutes.GET("/:id", func(ctx *gin.Context) {
				userController.Select(ctx)
			})

			userRoutes.GET("/all", func(ctx *gin.Context) {
				userController.SelectAll(ctx)
			})
		}

	}
}

func initController() {
	db := database.GetDB()
	userController = module.InitializeUserController(db)
	userRepository := repository.NewUserRepository(db)
	loginService := service.NewLoginService(userRepository)
	loginController = controller.NewLoginController(loginService)

}
