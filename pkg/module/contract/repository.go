package contract

import (
	"context"
	"time"

	"github.com/raymondsugiarto/funder-api/pkg/entity"
	"github.com/raymondsugiarto/funder-api/pkg/model"
	"github.com/raymondsugiarto/funder-api/shared/database/pagination"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, dto *entity.ContractDto) (*entity.ContractDto, error)
	FindByID(ctx context.Context, id string) (*entity.ContractDto, error)
	FindAll(ctx context.Context, req pagination.PaginationRequestDto) (*pagination.ResultPagination[entity.ContractDto], error)
	FindAllAging(ctx context.Context, req pagination.PaginationRequestDto) (*pagination.ResultPagination[entity.ContractDto], error)
	Update(ctx context.Context, dto *entity.ContractDto) (*entity.ContractDto, error)
	Delete(ctx context.Context, id string) error

	FindLastPerFunder(ctx context.Context, funderId string) (*entity.ContractDto, error)
	ViewDashboard(c context.Context, funderID string) (*entity.DashboardDto, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, dto *entity.ContractDto) (*entity.ContractDto, error) {
	m := dto.ToModel()
	err := r.db.Create(m).Error
	if err != nil {
		return nil, err
	}
	dto.ID = m.ID
	return dto, nil
}

func (r *repository) FindByID(ctx context.Context, id string) (*entity.ContractDto, error) {
	var m *model.Contract
	err := r.db.WithContext(ctx).First(&m, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return entity.NewContractDtoFromModel(m), nil
}

func (r *repository) FindLastPerFunder(ctx context.Context, funderId string) (*entity.ContractDto, error) {
	var m *model.Contract
	err := r.db.WithContext(ctx).First(&m, "funder_id = ?", funderId).Order("contract_number desc").Error
	if err != nil {
		return nil, err
	}
	return entity.NewContractDtoFromModel(m), nil
}

func (r *repository) FindAll(ctx context.Context, req pagination.PaginationRequestDto) (*pagination.ResultPagination[entity.ContractDto], error) {
	m := &model.Contract{}
	info, paginationResult, err := pagination.NewTable[*entity.ContractFilterDto, *model.Contract, entity.ContractDto]().
		Paginate(ctx, func(req *entity.ContractFilterDto) *gorm.DB {
			query := r.db.WithContext(ctx).Model(m).Preload("ContractPayments").Preload("Funder")
			// TODO: funder_id filter
			if req.FunderID != "" {
				query.Joins("INNER JOIN funder f ON f.id = contract.funder_id")
				query.Where("contract.funder_id = ? OR f.funder_id_parent = ?", req.FunderID, req.FunderID)
			}
			if req.NotYetPaidOff {
				query.Scopes(m.ScopeNotYetPaidOff)
			}

			return query
		}, &pagination.TableRequest[*entity.ContractFilterDto]{
			Request:       req.(*entity.ContractFilterDto),
			AllowedFields: []string{"funder_id"},
		})
	if err != nil {
		return nil, err
	}
	result := make([]entity.ContractDto, len(paginationResult))
	for i, m := range paginationResult {
		result[i] = new(entity.ContractDto).FromModel(m)
	}
	info.Data = result
	return info, nil
}

func (r *repository) FindAllAging(ctx context.Context, req pagination.PaginationRequestDto) (*pagination.ResultPagination[entity.ContractDto], error) {
	m := &model.Contract{}
	info, paginationResult, err := pagination.NewTable[*entity.ContractFilterDto, *model.Contract, entity.ContractDto]().
		Paginate(ctx, func(req *entity.ContractFilterDto) *gorm.DB {
			query := r.db.WithContext(ctx).
				Model(m).Preload("Funder")

			// yang akan overdue dan belum paid off
			query.Where("contract.due_date > ?", time.Now())
			query.Order("due_date asc")
			query.Scopes(m.ScopeNotYetPaidOff)
			return query
		}, &pagination.TableRequest[*entity.ContractFilterDto]{
			Request:       req.(*entity.ContractFilterDto),
			AllowedFields: []string{"funder_id"},
		})
	if err != nil {
		return nil, err
	}
	result := make([]entity.ContractDto, len(paginationResult))
	for i, m := range paginationResult {
		result[i] = new(entity.ContractDto).FromModel(m)
	}
	info.Data = result
	return info, nil
}

func (r *repository) Update(ctx context.Context, dto *entity.ContractDto) (*entity.ContractDto, error) {
	err := r.db.Save(dto.ToModel()).Error
	if err != nil {
		return nil, err
	}
	return dto, nil
}

func (r *repository) Delete(ctx context.Context, id string) error {
	err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Contract{}).Error
	if err != nil {
		return err
	}
	return err
}

func (r *repository) ViewDashboard(c context.Context, funderID string) (*entity.DashboardDto, error) {
	var m *entity.DashboardDto
	now := time.Now().Format(time.DateTime)
	query := r.db.WithContext(c).Table("contract").
		Select("SUM(amount) as total_amount, sum(return_amount) as total_return_amount, SUM(CASE WHEN disbursement_at <= '" + now + "' THEN amount ELSE 0 END) as total_amount_disbursed, SUM(CASE WHEN disbursement_at <='" + now + "' THEN return_amount ELSE 0 END) as total_return_amount_received, SUM(CASE WHEN disbursement_at >'" + now + "' THEN return_amount ELSE 0 END) as total_return_amount_pending")
	if funderID != "" {
		query.Joins("INNER JOIN funder f ON f.id = contract.funder_id")
		query.Where("contract.funder_id = ? OR f.funder_id_parent = ?", funderID, funderID)
	}

	err := query.Find(&m).Error
	if err != nil {
		return nil, err
	}
	return &entity.DashboardDto{
		TotalAmount:               m.TotalAmount,
		TotalAmountDisbursed:      m.TotalAmountDisbursed,
		TotalReturnAmount:         m.TotalReturnAmount,
		TotalReturnAmountReceived: m.TotalReturnAmountReceived,
		TotalReturnAmountPending:  m.TotalReturnAmountPending,
	}, nil
}
