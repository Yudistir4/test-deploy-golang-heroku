package users

import (
	"clean-code/app/middlewares"
	"fmt"
)

type UserUsecase struct {
	userRepository Repository
	jwtAuth        *middlewares.ConfigJwt
}

func NewUserUsecase(userRepository Repository, jwtAuth *middlewares.ConfigJwt) Usecase {

	return &UserUsecase{userRepository: userRepository, jwtAuth: jwtAuth}
}

func (uu *UserUsecase) CreateUser(domain *Domain) Domain {
	return uu.userRepository.CreateUser(domain)
}
func (uu *UserUsecase) Login(domain *Domain) string {
	user := uu.userRepository.GetByEmail(domain)

	fmt.Println(user)
	fmt.Println("Start")
	if user.ID == 0 {
		return ""
	}

	token := uu.jwtAuth.GenerateToken(int(user.ID))
	return token
}

func (uu *UserUsecase) GetAllUsers() []Domain {
	return uu.userRepository.GetAllUsers()
}
