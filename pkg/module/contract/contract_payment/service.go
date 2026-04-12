package contractpayment

import (
	"context"

	"github.com/gofiber/fiber/v3/log"
	"github.com/raymondsugiarto/funder-api/pkg/entity"
	"github.com/raymondsugiarto/funder-api/shared/database/pagination"
	"github.com/raymondsugiarto/funder-api/shared/database/transaction"
)

const ServiceName = "contractPaymentService"

type CallbackService interface {
	UpdateTotalPaidAmount(ctx context.Context, dto *entity.ContractPaymentDto) error
}

type Service interface {
	Create(ctx context.Context, dto *entity.ContractPaymentDto) (*entity.ContractPaymentDto, error)
	FindByID(ctx context.Context, id string) (*entity.ContractPaymentDto, error)
	FindAll(ctx context.Context, req pagination.PaginationRequestDto) (*pagination.ResultPagination[entity.ContractPaymentDto], error)
	Update(ctx context.Context, dto *entity.ContractPaymentDto) (*entity.ContractPaymentDto, error)
	Delete(ctx context.Context, id string) error
}

type service struct {
	txManager transaction.Manager
	repo      Repository
	callback  CallbackService
}

func NewService(
	txManager transaction.Manager,
	repo Repository,
	callback CallbackService,
) Service {
	return &service{
		txManager: txManager,
		repo:      repo,
		callback:  callback,
	}
}

func (s *service) Create(ctx context.Context, dto *entity.ContractPaymentDto) (*entity.ContractPaymentDto, error) {
	err := s.txManager.Execute(ctx, func(ctx context.Context) error {
		_, err := s.repo.Create(ctx, dto)
		if err != nil {
			return err
		}

		log.WithContext(ctx).Infof("update total paid amount %d", dto.PaymentAmount)
		// update total paid amount in contract
		err = s.callback.UpdateTotalPaidAmount(ctx, dto)
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	return dto, nil
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
