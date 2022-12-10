package repos

import (
	"github.com/ilmsg/golang-gorm-user-with-role/models"
	"gorm.io/gorm"
)

type userRepo struct {
	*gorm.DB
}

func (u *userRepo) Create(user *models.User) error {
	if err := u.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (*userRepo) Delete(id string) error {
	panic("unimplemented")
}

func (*userRepo) Get(id string) (*models.User, error) {
	panic("unimplemented")
}

func (u *userRepo) List() ([]*models.User, error) {
	var users []*models.User
	if err := u.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (*userRepo) Update(*models.User) error {
	panic("unimplemented")
}

func NewUserRepo(db *gorm.DB) models.UserRepo {
	return &userRepo{db}
}
