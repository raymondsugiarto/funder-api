package funder

import (
	"context"

	"github.com/raymondsugiarto/funder-api/pkg/entity"
	"github.com/raymondsugiarto/funder-api/pkg/module/user"
	"github.com/raymondsugiarto/funder-api/shared/database/pagination"
	"github.com/raymondsugiarto/funder-api/shared/database/transaction"
)

const ServiceName = "funderService"

type Service interface {
	Create(ctx context.Context, dto *entity.FunderDto) (*entity.FunderDto, error)
	FindByID(ctx context.Context, id string) (*entity.FunderDto, error)
	FindAll(ctx context.Context, req pagination.PaginationRequestDto) (*pagination.ResultPagination[entity.FunderDto], error)
	Update(ctx context.Context, dto *entity.FunderDto) (*entity.FunderDto, error)
	Delete(ctx context.Context, id string) error
}

type service struct {
	txManager transaction.Manager
	repo      Repository
	userSvc   user.Service
}

func NewService(
	txManager transaction.Manager,
	repo Repository,
	userSvc user.Service,
) Service {
	return &service{
		txManager: txManager,
		repo:      repo,
		userSvc:   userSvc,
	}
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

func (s *service) FindAll(ctx context.Context, req pagination.PaginationRequestDto) (*pagination.ResultPagination[entity.FunderDto], error) {
	return s.repo.FindAll(ctx, req)
}

func (s *service) Update(ctx context.Context, dto *entity.FunderDto) (*entity.FunderDto, error) {
	return s.repo.Update(ctx, dto)
}

func (s *service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
