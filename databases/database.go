package databases

import (
	"github.com/ilmsg/golang-gorm-user-with-role/models"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQLiteDB() *gorm.DB {
	var dbname string = "file::memory:?cache=shared"
	db, err := gorm.Open(sqlite.Open(dbname), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Role{})

	return db
}

func NewMySQLDB() *gorm.DB {
	var dns string = "root:test@tcp(127.0.0.1:3306)/user-with-role?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns))
	if err != nil {
		panic(err)
	}

	db.Migrator().DropTable(&models.User{})
	db.Migrator().DropTable(&models.Role{})

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Role{})

	return db
}
