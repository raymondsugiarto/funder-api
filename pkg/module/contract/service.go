package funder

import (
	"context"

	"github.com/raymondsugiarto/funder-api/pkg/entity"
	"github.com/raymondsugiarto/funder-api/shared/pagination"
)

const ServiceName = "funderService"

type Service interface {
	Create(ctx context.Context, dto *entity.ContractDto) (*entity.ContractDto, error)
	FindByID(ctx context.Context, id string) (*entity.ContractDto, error)
	FindAll(ctx context.Context, req pagination.PaginationRequestDto) (*pagination.ResultPagination[entity.ContractDto], error)
	Update(ctx context.Context, dto *entity.ContractDto) (*entity.ContractDto, error)
	Delete(ctx context.Context, id string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(ctx context.Context, dto *entity.ContractDto) (*entity.ContractDto, error) {
	return s.repo.Create(ctx, dto)
}

func (s *service) FindByID(ctx context.Context, id string) (*entity.ContractDto, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *service) FindAll(ctx context.Context, req pagination.PaginationRequestDto) (*pagination.ResultPagination[entity.ContractDto], error) {
	return s.repo.FindAll(ctx, req)
}

func (s *service) Update(ctx context.Context, dto *entity.ContractDto) (*entity.ContractDto, error) {
	return s.repo.Update(ctx, dto)
}

func (s *service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
