package repository

import (
	"auth-service/src/application"
	"auth-service/src/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	connection *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{connection: application.GlobalDB}
}

func (this *UserRepository) GetAll() ([]*model.User, error) {
	var users []*model.User
	err := this.connection.Find(&users).Error

	return users, err
}

func (this *UserRepository) GetById(id uint) (*model.User, error) {
	var user *model.User
	err := this.connection.First(&user, id).Error

	return user, err
}

func (this *UserRepository) GetByEmail(email string) (*model.User, error) {
	var user *model.User
	err := this.connection.Where("email = ?", email).First(&user).Error

	return user, err
}

func (this *UserRepository) Save(user model.User) (*model.User, error) {
	result := this.connection.Save(&user)

	return &user, result.Error
}

func (this *UserRepository) Delete(id uint) error {
	var user model.User
	if err := this.connection.First(&user, id).Error; err != nil {
		return err
	}

	return this.connection.Delete(&user).Error
}

func (this *UserRepository) HasPermission(id uint, permissionSlug string) bool {
	var count int64
	err := this.connection.
		Model(&model.Permission{}).
		Joins("JOIN roles_permissions ON permissions.id = roles_permissions.permission_id").
		Joins("JOIN roles ON roles.id = roles_permissions.role_id").
		Joins("JOIN users ON users.role_id = roles.id").
		Where("users.id = ? AND permissions.slug = ?", id, permissionSlug).
		Count(&count).Error

	if err != nil {
		return false
	}

	return count > 0
}
