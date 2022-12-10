package models

type User struct {
	ID       uint   `gorm:"primarykey"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleID   uint   `json:"-"`
	Role     *Role  `json:"Role" gorm:"foreignKey:RoleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (u *User) TableName() string {
	return "user"
}

type UserRepo interface {
	Get(id string) (*User, error)
	List() ([]*User, error)
	Create(*User) error
	Update(*User) error
	Delete(id string) error
}

type UserUsecase interface {
	Get(id string) (*User, error)
	List() ([]*User, error)
	Create(*User) error
	Update(*User) error
	Delete(id string) error
}
