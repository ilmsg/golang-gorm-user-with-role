package models

type Role struct {
	ID    uint   `gorm:"primarykey"`
	Name  string `json:"name"`
	Users []User `json:"users" gorm:"foreignKey:RoleID;"`
}

func (r *Role) TableName() string {
	return "role"
}

type RoleRepo interface {
	Get(id string) (*Role, error)
	List() ([]*Role, error)
	Create(*Role) error
	Update(*Role) error
	Delete(id string) error
}

type RoleUsecase interface {
	Get(id string) (*Role, error)
	List() ([]*Role, error)
	Create(*Role) error
	Update(*Role) error
	Delete(id string) error
}
