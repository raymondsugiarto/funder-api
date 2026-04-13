package funder

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/raymondsugiarto/funder-api/pkg/entity"
	"github.com/raymondsugiarto/funder-api/pkg/model"
	"github.com/raymondsugiarto/funder-api/pkg/module/user"
	usercredential "github.com/raymondsugiarto/funder-api/pkg/module/user-credential"
	"github.com/raymondsugiarto/funder-api/shared/database/pagination"
	"github.com/raymondsugiarto/funder-api/shared/database/transaction"
)

const ServiceName = "funderService"

type Service interface {
	Create(ctx context.Context, dto *entity.FunderDto) (*entity.FunderDto, error)
	FindByID(ctx context.Context, id string) (*entity.FunderDto, error)
	FindByUserID(ctx context.Context, userID string) (*entity.FunderDto, error)
	FindAll(ctx context.Context, req pagination.PaginationRequestDto) (*pagination.ResultPagination[entity.FunderDto], error)
	Update(ctx context.Context, dto *entity.FunderDto) (*entity.FunderDto, error)
	Delete(ctx context.Context, id string) error

	IdentifySessionFunder(ctx context.Context, userSession *entity.UserSessionDto) (*entity.FunderDto, error)
}

type service struct {
	txManager         transaction.Manager
	repo              Repository
	userSvc           user.Service
	userCredentialSvc usercredential.Service
}

func NewService(
	txManager transaction.Manager,
	repo Repository,
	userSvc user.Service,
	userCredentialSvc usercredential.Service,
) Service {
	return &service{
		txManager:         txManager,
		repo:              repo,
		userSvc:           userSvc,
		userCredentialSvc: userCredentialSvc,
	}
}

func (s *service) IdentifySessionFunder(ctx context.Context, userSession *entity.UserSessionDto) (*entity.FunderDto, error) {
	user := userSession.UserCredential.User
	if user.UserType == model.FUNDER {
		funder, err := s.FindByUserID(ctx, user.ID)
		if err != nil {
			log.WithContext(ctx).Errorf("error find funder by user id", err)
			return nil, err
		}
		return funder, nil
	}
	return nil, nil
}

func (s *service) Create(ctx context.Context, dto *entity.FunderDto) (*entity.FunderDto, error) {
	var funderDtoResult *entity.FunderDto
	err := s.txManager.Execute(ctx, func(txCtx context.Context) error {
		userDto, err := s.userSvc.Create(txCtx, dto.ToUserDto())
		if err != nil {
			return err
		}
		dto.UserID = userDto.ID
		res, err := s.repo.Create(txCtx, dto)
		if err != nil {
			return err
		}
		funderDtoResult = res
		return nil
	})
	return funderDtoResult, err
}

func (s *service) FindByID(ctx context.Context, id string) (*entity.FunderDto, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *service) FindByUserID(ctx context.Context, id string) (*entity.FunderDto, error) {
	return s.repo.FindByUserID(ctx, id)
}

func (s *service) FindAll(ctx context.Context, req pagination.PaginationRequestDto) (*pagination.ResultPagination[entity.FunderDto], error) {
	return s.repo.FindAll(ctx, req)
}

func (s *service) Update(ctx context.Context, newDto *entity.FunderDto) (*entity.FunderDto, error) {
	dto, err := s.FindByID(ctx, newDto.ID)
	if err != nil {
		return nil, err
	}
	dto.Name = newDto.Name
	dto.PhoneNumber = newDto.PhoneNumber
	dto.FunderIDParent = newDto.FunderIDParent

	err = s.txManager.Execute(ctx, func(txCtx context.Context) error {
		_, err := s.repo.Update(txCtx, dto)
		if err != nil {
			return err
		}
		if newDto.Password != "" {
			err = s.userCredentialSvc.ChangePassword(txCtx, dto.UserID, newDto.Password)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return dto, nil
}

func (s *service) Delete(ctx context.Context, id string) error {
	// TODO: validation have contract and child
	return s.repo.Delete(ctx, id)
}
