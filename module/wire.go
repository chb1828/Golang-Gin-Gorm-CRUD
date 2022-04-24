//go:build wireinject
// +build wireinject

package module

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"login/controller"
	"login/core/repository"
	"login/service"
)

func InitializeUserController(db *gorm.DB) controller.UserController {
	wire.Build(repository.NewUserRepository, service.NewUserService, controller.NewUserController)
	return controller.InUserController{}
}

/*
func InitializeLoginController(db *gorm.DB) controller.LoginController {
	wire.Build(repository.NewUserRepository, service.NewLoginService, controller.NewLoginController)
	return controller.InLoginController{}
}
*/
