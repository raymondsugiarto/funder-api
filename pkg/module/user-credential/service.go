package usercredential

import (
	"context"

	"github.com/raymondsugiarto/funder-api/pkg/entity"
	"github.com/raymondsugiarto/funder-api/pkg/model"
	"github.com/raymondsugiarto/funder-api/shared/security"
)

const ServiceName = "userCredentialService"

type Service interface {
	FindByUsername(ctx context.Context, userCredential *entity.UserCredentialDto) (*entity.UserCredentialDto, error)
	FindByEmail(ctx context.Context, userCredential *entity.UserCredentialDto) (*entity.UserCredentialDto, error)
	GetUserCredentialByUsername(ctx context.Context, organizationID, u string) (*model.UserCredential, error)
	ChangePassword(ctx context.Context, userID, password string) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

// FindByUsername is a function to find user credential by username
func (s *service) FindByUsername(ctx context.Context, userCredential *entity.UserCredentialDto) (*entity.UserCredentialDto, error) {
	return s.repository.FindByUsername(ctx, userCredential)
}

// FindByEmail is a function to find user credential by username
func (s *service) FindByEmail(ctx context.Context, userCredential *entity.UserCredentialDto) (*entity.UserCredentialDto, error) {
	return s.repository.FindByEmail(ctx, userCredential)
}

func (s *service) GetUserCredentialByUsername(ctx context.Context, organizationID, u string) (*model.UserCredential, error) {
	return s.repository.GetUserCredentialByUsername(ctx, organizationID, u)
}

func (s *service) ChangePassword(ctx context.Context, userID, password string) error {
	encryptedPass, _ := security.HashPassword(password)
	return s.repository.ChangePassword(ctx, userID, encryptedPass)
}
