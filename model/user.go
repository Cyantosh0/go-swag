package model

import (
	"Cyantosh0/go-swag/config"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID         BinaryUUID     `json:"id"`
	Name       string         `json:"name"`
	Email      string         `json:"email"`
	Phone      string         `json:"phone"`
	Address    string         `json:"address"`
	UserNumber sql.NullString `json:"user_number" swaggertype:"string"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	id, err := uuid.NewRandom()
	u.ID = BinaryUUID(id)
	return err
}

//GetAllUsers Fetch all user data
func GetAllUsers(user *[]User) error {
	if err := config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

//CreateUser ... Insert New data
func CreateUser(user *User) error {
	if err := config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByID ... Fetch only one user by Id
func GetUserByID(user *User, id BinaryUUID) error {
	if err := config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}
