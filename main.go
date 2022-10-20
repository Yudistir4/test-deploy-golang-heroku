package main

import (
	"clean-code/app/middlewares"
	"clean-code/routes"
	"log"

	_userUseCase "clean-code/businesses/users"
	_userController "clean-code/controllers/users"
	"clean-code/drivers"
	"clean-code/drivers/mysql"
	"clean-code/util"

	"github.com/labstack/echo/v4"
)

func main() {
	configDB := mysql.ConfigDB{
		DB_USERNAME: util.GetConfig("DB_USERNAME"),
		DB_PASSWORD: util.GetConfig("DB_PASSWORD"),
		DB_HOST:     util.GetConfig("DB_HOST"),
		DB_PORT:     util.GetConfig("DB_PORT"),
		DB_NAME:     util.GetConfig("DB_NAME"),
	}

	db := configDB.InitDB()

	mysql.DBMigrate(db)

	configJWT := middlewares.ConfigJwt{
		SecretJWT:       util.GetConfig("JWT_SECRET_KEY"),
		ExpiresDuration: 1,
	}

	e := echo.New()

	userRepo := drivers.NewUserRepository(db)
	userUsecase := _userUseCase.NewUserUsecase(userRepo, &configJWT)
	userController := _userController.NewAuthController(userUsecase)

	routesInit := routes.ControllerList{
		JWTMiddleware:  configJWT.Init(),
		AuthController: *userController,
	}

	routesInit.RouteRegister(e)

	log.Fatal(e.Start(":8000"))
}