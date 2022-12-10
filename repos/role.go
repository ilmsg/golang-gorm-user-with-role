package repos

import (
	"github.com/ilmsg/golang-gorm-user-with-role/models"
	"gorm.io/gorm"
)

type roleRepo struct {
	*gorm.DB
}

func (r *roleRepo) Create(role *models.Role) error {
	if err := r.DB.Create(role).Error; err != nil {
		return err
	}
	return nil
}

func (*roleRepo) Delete(id string) error {
	panic("unimplemented")
}

func (*roleRepo) Get(id string) (*models.Role, error) {
	panic("unimplemented")
}

func (r *roleRepo) List() ([]*models.Role, error) {
	var roles []*models.Role
	if err := r.DB.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (*roleRepo) Update(*models.Role) error {
	panic("unimplemented")
}

func NewRoleRepo(db *gorm.DB) models.RoleRepo {
	return &roleRepo{db}
}
