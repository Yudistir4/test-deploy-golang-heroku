package users_test

import (
	"clean-code/app/middlewares"
	"clean-code/businesses/users"
	_userMock "clean-code/businesses/users/mocks"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	usersRepository _userMock.Repository
	usersService    users.Usecase

	usersDomain users.Domain
)

func TestMain(m *testing.M) {
	usersService = users.NewUserUsecase(&usersRepository, &middlewares.ConfigJwt{})

	usersDomain = users.Domain{
		Email:    "user1@mail.com",
		Password: "123456",
	}

	m.Run()
}

func TestCreateUser(t *testing.T) {
	t.Run("CreateUser | Valid", func(t *testing.T) {
		usersRepository.On("CreateUser", &usersDomain).Return(usersDomain).Once()

		result := usersService.CreateUser(&usersDomain)

		assert.NotNil(t, result)
	})

	t.Run("CreateUser | InValid", func(t *testing.T) {
		usersRepository.On("CreateUser", &users.Domain{}).Return(users.Domain{}).Once()

		result := usersService.CreateUser(&users.Domain{})

		assert.NotNil(t, result)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Login | Valid", func(t *testing.T) {
		usersRepository.On("GetByEmail", &usersDomain).Return(users.Domain{}).Once()

		result := usersService.Login(&usersDomain)

		assert.NotNil(t, result)
	})

	t.Run("Login | InValid", func(t *testing.T) {
		usersRepository.On("GetByEmail", &users.Domain{}).Return(users.Domain{}).Once()

		result := usersService.Login(&users.Domain{})

		assert.Empty(t, result)
	})

}

func TestGetAllUsers(t *testing.T) {
	t.Run("Get All | Valid", func(t *testing.T) {
		usersRepository.On("GetAllUsers").Return([]users.Domain{usersDomain}).Once()

		result := usersService.GetAllUsers()

		fmt.Println("result", result)
		assert.Equal(t, 1, len(result))
	})

}
