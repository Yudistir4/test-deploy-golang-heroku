package users

import (
	"clean-code/businesses/users"
	"clean-code/controllers"
	"clean-code/controllers/users/request"
	"clean-code/controllers/users/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authUsecase users.Usecase
}

func NewAuthController(authUC users.Usecase) *AuthController {
	return &AuthController{authUsecase: authUC}
}

func (ctrl *AuthController) CreateUser(c echo.Context) error {

	userInput := request.User{}

	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	err := userInput.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "validation failed",
		})

	}

	user := ctrl.authUsecase.CreateUser(userInput.ToDomain())
	return c.JSON(http.StatusCreated, response.FromDomain(user))
}
func (ctrl *AuthController) Login(c echo.Context) error {
	userInput := request.User{}

	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	err := userInput.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "validation failed",
		})

	}

	token := ctrl.authUsecase.Login(userInput.ToDomain())

	if token == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "invalid email or password",
		})

	}
	return c.JSON(http.StatusCreated, map[string]string{
		"token": token,
	})

}
func (ctrl *AuthController) GetAllUsers(c echo.Context) error {

	usersData := ctrl.authUsecase.GetAllUsers()

	users := []response.User{}

	for _, user := range usersData {
		users = append(users, response.FromDomain(user))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all users", users)

}
