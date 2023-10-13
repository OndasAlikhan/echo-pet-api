package service

import (
	"auth-service/src/model"
	"auth-service/src/model/dto"
	"auth-service/src/repository"
	"github.com/davecgh/go-spew/spew"
)

type UserService struct {
	repository  *repository.UserRepository
	roleService *RoleService
	jwtService  *JWTService
}

func NewUserService() *UserService {
	return &UserService{
		repository:  repository.NewUserRepository(),
		roleService: NewRoleService(),
		jwtService:  NewJWTService(),
	}
}

func (this *UserService) GetAll() ([]*model.User, error) {
	return this.repository.GetAll()
}

func (this *UserService) GetById(id uint) (*model.User, error) {
	return this.repository.GetById(id)
}

func (this *UserService) GetByEmail(email string) (*model.User, error) {
	return this.repository.GetByEmail(email)
}

func (this *UserService) Create(request *dto.UserStoreRequest) (*model.User, error) {
	passwordHash, err := this.jwtService.GenerateHash(request.Password)
	if err != nil {
		return nil, err
	}

	spew.Dump(request)

	role, err := this.roleService.GetBySlug("user")
	if err != nil {
		return nil, err
	}

	user := model.User{Email: request.Email, Name: request.Name, Password: passwordHash, RoleID: role.ID}

	return this.repository.Save(user)
}

func (this *UserService) UpdatePassword(id uint, request *dto.UserUpdatePasswordRequest) (*model.User, error) {
	user, err := this.repository.GetById(id)
	if err != nil {
		return nil, err
	}

	user.Password, err = this.jwtService.GenerateHash(request.Password)
	if err != nil {
		return nil, err
	}

	return this.repository.Save(*user)
}

func (this *UserService) Delete(id uint) error {
	return this.repository.Delete(id)
}

func (this *UserService) HasPermission(id uint, permissionSlug string) bool {
	return this.repository.HasPermission(id, permissionSlug)
}
