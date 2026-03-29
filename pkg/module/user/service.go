package user

import (
	"context"
	"errors"

	"github.com/raymondsugiarto/funder-api/pkg/entity"
	usercredential "github.com/raymondsugiarto/funder-api/pkg/module/user-credential"
)

type Service interface {
	CreateUser(*entity.CreateUser) (*entity.CreateUser, error)
	FindByUserID(ctx context.Context, userID string) (*entity.UserDto, error)
}

type service struct {
	repository            Repository
	userCredentialService usercredential.Service
}

func NewService(repository Repository, userCredentialService usercredential.Service) Service {
	return &service{
		repository:            repository,
		userCredentialService: userCredentialService,
	}
}

// FindByUserID is a function to find user by user id
func (s *service) FindByUserID(ctx context.Context, userID string) (*entity.UserDto, error) {
	userDto, err := s.repository.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return userDto, nil
}

// CreateUser is a function to create user
func (s *service) CreateUser(createUser *entity.CreateUser) (*entity.CreateUser, error) {
	userCredential := &entity.UserCredential{
		OrganizationData: createUser.OrganizationData,
		Username:         createUser.Username,
	}
	_, err := s.userCredentialService.FindByUsername(userCredential)
	if err == nil {
		return nil, errors.New("errorAccountCodeAlreadyExist")
	}

	_, err = s.userCredentialService.FindByEmail(userCredential)
	if err == nil {
		return nil, errors.New("errorEmailAlreadyExist")
	}

	return s.repository.CreateUser(createUser)
}
