package contract

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v2/log"
	"github.com/raymondsugiarto/funder-api/pkg/entity"
	"github.com/raymondsugiarto/funder-api/shared/database/pagination"
	"github.com/raymondsugiarto/funder-api/shared/database/transaction"
	"github.com/raymondsugiarto/funder-api/shared/database/view"
	"gorm.io/gorm"
)

const ServiceName = "contractService"

const (
	ViewList  = "list"
	ViewAging = "aging"
)

type ViewFunc func(ctx context.Context, req pagination.PaginationRequestDto) (*pagination.ResultPagination[entity.ContractDto], error)

type Service interface {
	DashboardService
	Create(ctx context.Context, dto *entity.ContractDto) (*entity.ContractDto, error)
	FindByID(ctx context.Context, id string) (*entity.ContractDto, error)
	FindAll(ctx context.Context, req pagination.PaginationRequestDto) (*pagination.ResultPagination[entity.ContractDto], error)
	Update(ctx context.Context, dto *entity.ContractDto) (*entity.ContractDto, error)
	Delete(ctx context.Context, id string) error

	UpdateTotalPaidAmount(ctx context.Context, dto *entity.ContractPaymentDto) error
}

type service struct {
	txManager transaction.Manager
	view      view.Service[ViewFunc]
	repo      Repository
}

func NewService(
	txManager transaction.Manager, repo Repository,
) Service {
	viewService := view.NewViewService[ViewFunc]()
	s := &service{
		txManager: txManager,
		view:      viewService,
		repo:      repo,
	}
	viewService.Add(ViewList, s.ViewList)
	viewService.Add(ViewAging, s.ViewAging)
	return s
}

func (s *service) Create(ctx context.Context, dto *entity.ContractDto) (*entity.ContractDto, error) {
	dto.ReturnAmount = dto.Amount * dto.ReturnPercentage / 100
	lastContract, err := s.repo.FindLastPerFunder(ctx, dto.FunderID)
	if err != nil {
		// check if error is not record not found, then return error, otherwise continue with contract number 1
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		lastContract = entity.NewContractDto()
		lastContract.ContractNumber = 0
	}
	dto.ContractNumber = lastContract.ContractNumber + 1
	return s.repo.Create(ctx, dto)
}

func (s *service) UpdateTotalPaidAmount(ctx context.Context, dto *entity.ContractPaymentDto) error {
	contract, err := s.repo.FindByID(ctx, dto.ContractID)
	if err != nil {
		log.WithContext(ctx).Errorf("error find contract %v", err)
		return err
	}
	contract.TotalPaidAmount += dto.PaymentAmount
	_, err = s.repo.Update(ctx, contract)
	if err != nil {
		log.WithContext(ctx).Errorf("error update contract %v", err)
		return err
	}
	return nil
}

func (s *service) FindByID(ctx context.Context, id string) (*entity.ContractDto, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *service) FindAll(ctx context.Context, req pagination.PaginationRequestDto) (*pagination.ResultPagination[entity.ContractDto], error) {
	return s.view.Get(req.GetView())(ctx, req)
}

func (s *service) ViewList(ctx context.Context, req pagination.PaginationRequestDto) (*pagination.ResultPagination[entity.ContractDto], error) {
	return s.repo.FindAll(ctx, req)
}

func (s *service) ViewAging(ctx context.Context, req pagination.PaginationRequestDto) (*pagination.ResultPagination[entity.ContractDto], error) {
	return s.repo.FindAllAging(ctx, req)
}

func (s *service) Update(ctx context.Context, newDto *entity.ContractDto) (*entity.ContractDto, error) {
	dto, err := s.FindByID(ctx, newDto.ID)
	if err != nil {
		return nil, err
	}

	dto.ContractCode = newDto.ContractCode
	dto.FunderID = newDto.FunderID
	dto.Amount = newDto.Amount
	dto.ReturnPercentage = newDto.ReturnPercentage
	dto.Duration = newDto.Duration
	dto.Notes = newDto.Notes
	dto.DestinationAccount = newDto.DestinationAccount
	if newDto.AttachmentURL != "" {
		dto.AttachmentURL = newDto.AttachmentURL
	}

	return s.repo.Update(ctx, dto)
}

func (s *service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
