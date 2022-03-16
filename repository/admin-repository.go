package repository

import (
	"log"

	"github.com/doffy007/price-comparison-api.git/entities"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AdminRepository interface {
	InsertAdmin(admin entities.Admin) entities.Admin
	UpdateAdmin(admin entities.Admin) entities.Admin
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) entities.Admin
	ProfileAdmin(adminID string) entities.Admin
}

type adminConnection struct {
	connection *gorm.DB
}

func NewAdminReposistory(db *gorm.DB) AdminRepository {
	return &adminConnection{
		connection: db,
	}
}

func (db *adminConnection) InsertAdmin(admin entities.Admin) entities.Admin {
	admin.Password = HashAndSalt([]byte(admin.Password))
	db.connection.Save(&admin)
	return admin
}

func (db *adminConnection) UpdateAdmin(admin entities.Admin) entities.Admin {
	admin.Password =  HashAndSalt([]byte(admin.Password))
	db.connection.Save(&admin)
	return admin
}

func (db *adminConnection) VerifyCredential(email string, password string) interface{} {
	var admin entities.Admin
	res := db.connection.Where("email = ?", email).Take(&admin)
	if res.Error == nil {
		return admin
	}
	return nil
}

func (db *adminConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var admin entities.Admin
	return db.connection.Where("email = ?", email).Take(&admin)
}

func (db *adminConnection) FindByEmail(email string) entities.Admin {
	var admin entities.Admin
	db.connection.Where("email = ?", email).Take(&admin)
	return admin
}

func (db *adminConnection) ProfileAdmin(adminID string) entities.Admin {
	var admin entities.Admin
	db.connection.Find(&admin, adminID)
	return admin
}

func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("cannot hash a password ")
	}
	return string(hash)
}
