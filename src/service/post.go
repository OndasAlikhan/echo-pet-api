package service

import (
	"echo-pet-api/src/model"
	"echo-pet-api/src/model/dto"
	"echo-pet-api/src/repository"
	"github.com/gosimple/slug"
)

type PostService struct {
	repository *repository.PostRepository
}

func NewPostService() *PostService {
	return &PostService{repository: repository.NewPostRepository()}
}

func (ps *PostService) GetAll() ([]*dto.PostResponse, error) {
	posts, err := ps.repository.GetAll()
	if err != nil {
		return nil, err
	}

	response := []*dto.PostResponse{}
	for _, post := range posts {
		response = append(response, dto.NewPostResponse(post))
	}

	return response, nil
}

func (ps *PostService) GetById(id uint) (*dto.PostResponse, error) {
	post, err := ps.repository.GetById(id)
	if err != nil {
		return nil, err
	}

	return dto.NewPostResponse(post), nil
}

func (ps *PostService) GetBySlug(slug string) (*dto.PostResponse, error) {
	post, err := ps.repository.GetBySlug(slug)
	if err != nil {
		return nil, err
	}

	return dto.NewPostResponse(post), nil
}

func (ps *PostService) Create(request *dto.PostStoreRequest) (*dto.PostResponse, error) {
	if request.Slug == "" {
		request.Slug = slug.Make(request.Title)
	}

	post := model.Post{Slug: request.Slug, Title: request.Title, Desc: request.Desc}

	err := ps.repository.Create(&post)
	if err != nil {
		return nil, err
	}

	return dto.NewPostResponse(&post), nil
}

func (ps *PostService) Update(id uint, request *dto.PostUpdateRequest) (*dto.PostResponse, error) {
	post, err := ps.repository.GetById(id)
	if err != nil {
		return nil, err
	}

	post.Desc = request.Desc
	if err = ps.repository.Update(post); err != nil {
		return nil, err
	}

	return dto.NewPostResponse(post), nil
}

func (ps *PostService) Delete(id uint) error {
	post, err := ps.repository.GetById(id)
	if err != nil {
		return err
	}

	err = ps.repository.Delete(post)
	if err != nil {
		return err
	}

	return nil
}
