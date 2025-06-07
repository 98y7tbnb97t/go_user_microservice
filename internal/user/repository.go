package user

import (
	db "github.com/98y7tbnb97t/users-service/internal/database"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAllUsersFromBd() ([]User, error) {
	var users []User
	result := db.DB.Omit("created_at", "updated_at", "deleted_at").Find(&users)
	return users, result.Error
}

func (r *Repository) CreateUserFromBd(user *User) error {
	return db.DB.Create(user).Error
}

func (r *Repository) GetUserByIDFromBd(id int) (*User, error) {
	var user User
	if err := db.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) DeleteUserFromBd(id int) error {
	return db.DB.Delete(&User{}, id).Error
}

func (r *Repository) UpdateUserFromBd(id int, updates *User) error {
	var user User
	if err := db.DB.First(&user, id).Error; err != nil {
		return err
	}
	return db.DB.Model(&user).Updates(updates).Error
}
