package service

import (
	"auth-service/src/model"
	"auth-service/src/model/dto"
	"auth-service/src/repository"
	"github.com/gosimple/slug"
)

type PermissionService struct {
	repository *repository.PermissionRepository
}

func NewPermissionService() *PermissionService {
	return &PermissionService{repository: repository.NewPermissionRepository()}
}

func (this *PermissionService) GetAllByRole(roleID uint) ([]*model.Permission, error) {
	return this.repository.GetAllByRole(roleID)
}

func (this *PermissionService) GetById(id uint) (*model.Permission, error) {
	return this.repository.GetById(id)
}

func (this *PermissionService) Create(request *dto.PermissionStoreRequest) (*model.Permission, error) {
	permission := model.Permission{Name: request.Name, Slug: slug.Make(request.Name)}

	return this.repository.Save(permission)
}

func (this *PermissionService) Update(id uint, request *dto.PermissionUpdateRequest) (*model.Permission, error) {
	permission, err := this.repository.GetById(id)
	if err != nil {
		return nil, err
	}

	permission.Name = request.Name
	permission.Slug = slug.Make(request.Name)

	return this.repository.Save(*permission)
}

func (this *PermissionService) Delete(id uint) error {
	return this.repository.Delete(id)
}
