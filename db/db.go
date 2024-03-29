package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ashishthakur913/project/model"
	"os"
)

func New() *gorm.DB {
	db, err := gorm.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/database?parseTime=True")
	if err != nil {
		fmt.Println("storage err: ", err)
	}
	fmt.Println("ASHISH")
	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
	return db
}

func TestDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./../realworld_test.db")
	if err != nil {
		fmt.Println("storage err: ", err)
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(false)
	return db
}

func DropTestDB() error {
	if err := os.Remove("./../realworld_test.db"); err != nil {
		return err
	}
	return nil
}

//TODO: err check
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
		&model.Follow{},
		&model.Article{},
		&model.Comment{},
		&model.Tag{},
	)
}
