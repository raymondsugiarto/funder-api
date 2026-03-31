package funder

import (
	"context"

	"github.com/raymondsugiarto/funder-api/pkg/entity"
	"github.com/raymondsugiarto/funder-api/shared/database/pagination"
)

const ServiceName = "funderService"

type Service interface {
	Create(ctx context.Context, dto *entity.ContractPaymentDto) (*entity.ContractPaymentDto, error)
	FindByID(ctx context.Context, id string) (*entity.ContractPaymentDto, error)
	FindAll(ctx context.Context, req pagination.PaginationRequestDto) (*pagination.ResultPagination[entity.ContractPaymentDto], error)
	Update(ctx context.Context, dto *entity.ContractPaymentDto) (*entity.ContractPaymentDto, error)
	Delete(ctx context.Context, id string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(ctx context.Context, dto *entity.ContractPaymentDto) (*entity.ContractPaymentDto, error) {
	return s.repo.Create(ctx, dto)
}

func (s *service) FindByID(ctx context.Context, id string) (*entity.ContractPaymentDto, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *service) FindAll(ctx context.Context, req pagination.PaginationRequestDto) (*pagination.ResultPagination[entity.ContractPaymentDto], error) {
	return s.repo.FindAll(ctx, req)
}

func (s *service) Update(ctx context.Context, dto *entity.ContractPaymentDto) (*entity.ContractPaymentDto, error) {
	return s.repo.Update(ctx, dto)
}

func (s *service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
