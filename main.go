package main

import (
	"fmt"

	"github.com/ilmsg/golang-gorm-user-with-role/databases"
	"github.com/ilmsg/golang-gorm-user-with-role/models"
	"github.com/ilmsg/golang-gorm-user-with-role/repos"
)

func main() {

	// db := databases.NewSQLiteDB()
	db := databases.NewMySQLDB()

	roleRepo := repos.NewRoleRepo(db)

	roleAdmin := models.Role{Name: "Administator"}
	roleMember := models.Role{Name: "Member"}
	roleRepo.Create(&roleAdmin)
	roleRepo.Create(&roleMember)

	roles, _ := roleRepo.List()
	for _, role := range roles {
		fmt.Println(role)
	}

	userRepo := repos.NewUserRepo(db)

	userAdmin1 := models.User{Role: &roleAdmin, Email: "admin1@example.com"}
	userAdmin2 := models.User{Role: &roleAdmin, Email: "admin2@example.com"}
	userMember1 := models.User{Role: &roleMember, Email: "member1@example.com"}
	userMember2 := models.User{Role: &roleMember, Email: "member2@example.com"}
	userRepo.Create(&userAdmin1)
	userRepo.Create(&userAdmin2)
	userRepo.Create(&userMember1)
	userRepo.Create(&userMember2)

	users, _ := userRepo.List()
	for _, user := range users {
		fmt.Println(user)
	}
}
