package repository

import (
	"go-web-template/modules/model"
	"go-web-template/modules/orm/mysql"

	"github.com/google/wire"
	"gorm.io/gorm"
)

type IUserRepository interface {
	FindByEmail(email string) (user model.User, tx *gorm.DB)
	Create(user model.User) error
}

type UserRepository struct {
	MySQLGorm mysql.IMySQLGorm
}

var userRepositorySet = wire.NewSet(
	wire.Bind(new(IUserRepository), new(UserRepository)),
	UserRepositoryProvider,
)

func UserRepositoryProvider(mysqlGorm mysql.IMySQLGorm) UserRepository {
	return UserRepository{
		MySQLGorm: mysqlGorm,
	}
}

func (ur UserRepository) FindByEmail(email string) (user model.User, tx *gorm.DB) {
	tx = ur.MySQLGorm.Get().Where("email = ?", email).Find(&user)
	return user, tx
}

func (ur UserRepository) Create(user model.User) error {
	return ur.MySQLGorm.Get().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		return nil
	})
}
