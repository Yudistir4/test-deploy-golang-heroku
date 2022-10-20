package drivers

import (
	"clean-code/businesses/users"
	usersDB "clean-code/drivers/mysql/users"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) users.Repository {
	return usersDB.NewMySQLRepository(conn)
}
