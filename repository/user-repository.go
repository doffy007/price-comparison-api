package repository

import (
	"github.com/doffy007/price-comparison-api.git/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user entities.User) entities.User
	UpdateUser(user entities.User) entities.User
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) entities.User
	ProfileUser(userID string) entities.User
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) InsertUser(user entities.User) entities.User {
	user.Password = HashAndSalt([]byte(user.Password))
	db.connection.Save(&user)
	return user
}

func (db *userConnection) UpdateUser(user entities.User) entities.User {
	user.Password = HashAndSalt([]byte(user.Password))
	db.connection.Save(&user)
	return user
}

func (db *userConnection) VerifyCredential(email string, pasword string) interface{} {
	var user entities.User
	result := db.connection.Where("email = ?", email).Take(&user)
	if result.Error == nil {
		return user
	}
	return nil
}

func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entities.Admin
	return db.connection.Where("email = ?", email).Take(&user)
}

func (db *userConnection) FindByEmail(email string) entities.User {
	var user entities.User
	db.connection.Where("email = ?", email).Take(&user)
	return user
}

func (db *userConnection) ProfileUser(userID string) entities.User {
	var user entities.User
	db.connection.Find(&user, userID)
	return user
}
