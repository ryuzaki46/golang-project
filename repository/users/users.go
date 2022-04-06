package user

import (
	_entities "group-project/limamart/entities"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func HashPassword(password string) (string, error) {
	// Ini gue bikin 32 bytes format, kalau mau ganti angka 32 nya sesuai kebutuhan
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 32)
	return string(pass), err
}

func (ur *UserRepository) CreateUser(request _entities.User) (_entities.User, error) {
	fmt.Println(&request)
	// Tinggal panggil HashPassword tuk hashing nya. Sebelum masukkin ke DB
	yx := ur.DB.Save(&request)
	if yx.Error != nil {
		return request, yx.Error
	}

	return request, nil
}

func (ur *UserRepository) GetUserById(id int) (_entities.User, int, error) {
	var users _entities.User
	tx := ur.DB.Find(&users, id)
	if tx.Error != nil {
		return users, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return users, 0, nil
	}
	return users, int(tx.RowsAffected), nil
}

func (ur *UserRepository) UpdateUser(request _entities.User) (_entities.User, int, error) {
	tx := ur.DB.Save(&request)
	if tx.Error != nil {
		return request, 0, tx.Error
	}
	return request, int(tx.RowsAffected), nil
}

func (ur *UserRepository) DeleteUser(id int) error {

	err := ur.DB.Unscoped().Delete(&_entities.User{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
