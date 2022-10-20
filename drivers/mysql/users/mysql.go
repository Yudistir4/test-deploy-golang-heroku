package users

import (
	"clean-code/businesses/users"

	"gorm.io/gorm"
)

type UserRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) users.Repository {
	return &UserRepository{
		conn: conn,
	}
}

func (ur *UserRepository) CreateUser(domain *users.Domain) users.Domain {
	user := FromDomain(domain)
	ur.conn.Save(&user)

	return user.ToDomain()
}
func (ur *UserRepository) GetByEmail(domain *users.Domain) users.Domain {
	var user User
	ur.conn.First(&user, "email = ?", domain.Email)

	if user.ID == 0 {
		return users.Domain{}
	}

	if user.Password != domain.Password {
		return users.Domain{}
	}
	return user.ToDomain()
}
func (ur *UserRepository) GetAllUsers() []users.Domain {
	var rec []User
	ur.conn.Find(&rec)

	usersDomain := []users.Domain{}

	for _, user := range rec {
		usersDomain = append(usersDomain, user.ToDomain())
	}

	return usersDomain
}
