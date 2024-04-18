package store

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type UserRecord struct {
	gorm.Model
	ChatId   string
	UserId   string
	Message  string
	IsBanned bool
}

func Initdb(dsn string) (*gorm.DB, error) {
	sqldb, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(dsn, sqldb)
	if err != nil {
		fmt.Println("couldnt connect to db r could not find test db file")
	}

	db.AutoMigrate(&UserRecord{})

	return db, nil
}

func CreateUserRecord(db *gorm.DB, chatId string, messageId string, userId string) ([]UserRecord, error) {
	result := db.Create(&UserRecord{
		Model:    gorm.Model{},
		ChatId:   chatId,
		Message:  messageId,
		IsBanned: false,
	})
}

func GetUserRecord(db *gorm.DB, chatId string, messageId string, userId string) ([]UserRecord, error) {
	var users []UserRecord

	result := db.Where("ChatId = ? AND UserId = ?", chatId, userId).Find(&users)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
	}
	if len(users) == 0 {
		return nil, errors.New("no users found here")
	}
	return users, nil
}
