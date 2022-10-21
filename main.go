package main

import (
	"clean-code/app/middlewares"
	"clean-code/routes"
	"log"
	"os"

	_userUseCase "clean-code/businesses/users"
	_userController "clean-code/controllers/users"
	"clean-code/drivers"
	"clean-code/drivers/mysql"

	"github.com/labstack/echo/v4"
)

func main() {
	configDB := mysql.ConfigDB{
		// DB_USERNAME: util.GetConfig("DB_USERNAME"),
		// DB_PASSWORD: util.GetConfig("DB_PASSWORD"),
		// DB_HOST:     util.GetConfig("DB_HOST"),
		// DB_PORT:     util.GetConfig("DB_PORT"),
		// DB_NAME:     util.GetConfig("DB_NAME"),
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_NAME:     os.Getenv("DB_NAME"),
	}

	db := configDB.InitDB()

	mysql.DBMigrate(db)

	configJWT := middlewares.ConfigJwt{
		// SecretJWT:       util.GetConfig("JWT_SECRET_KEY"),
		SecretJWT:       os.Getenv("JWT_SECRET_KEY"),
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
