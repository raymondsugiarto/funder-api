package contractpayment

import (
	"context"

	"github.com/raymondsugiarto/funder-api/pkg/entity"
	"github.com/raymondsugiarto/funder-api/pkg/model"
	"github.com/raymondsugiarto/funder-api/shared/database/pagination"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, dto *entity.ContractPaymentDto) (*entity.ContractPaymentDto, error)
	FindByID(ctx context.Context, id string) (*entity.ContractPaymentDto, error)
	FindAll(ctx context.Context, req pagination.PaginationRequestDto) (*pagination.ResultPagination[entity.ContractPaymentDto], error)
	Update(ctx context.Context, dto *entity.ContractPaymentDto) (*entity.ContractPaymentDto, error)
	Delete(ctx context.Context, id string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, dto *entity.ContractPaymentDto) (*entity.ContractPaymentDto, error) {
	m := dto.ToModel()
	err := r.db.Create(m).Error
	if err != nil {
		return nil, err
	}
	dto.ID = m.ID
	return dto, nil
}

func (r *repository) FindByID(ctx context.Context, id string) (*entity.ContractPaymentDto, error) {
	var m entity.ContractPaymentDto
	err := r.db.WithContext(ctx).First(&m, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *repository) FindAll(ctx context.Context, req pagination.PaginationRequestDto) (*pagination.ResultPagination[entity.ContractPaymentDto], error) {
	info, paginationResult, err := pagination.NewTable[*entity.ContractPaymentFilterDto, *model.ContractPayment, entity.ContractPaymentDto]().
		Paginate(ctx, func(req *entity.ContractPaymentFilterDto) *gorm.DB {
			query := r.db.WithContext(ctx).Model(&model.ContractPayment{}).
				Joins("INNER JOIN contract on contract.id = contract_payment.contract_id").
				Preload("Contract").Preload("Contract.Funder")
			if req.FunderID != "" {
				query.Joins("INNER JOIN funder f ON f.id = contract.funder_id")
				query.Where("contract.funder_id = ? OR f.funder_id_parent = ?", req.FunderID, req.FunderID)
			}
			return query
		}, &pagination.TableRequest[*entity.ContractPaymentFilterDto]{
			Request:       req.(*entity.ContractPaymentFilterDto),
			QueryField:    []string{"notes"},
			AllowedFields: []string{"funder_id", "payment_at"},
			MapFields: map[string]string{
				"funder_id": "contract.funder_id",
			},
		})
	if err != nil {
		return nil, err
	}
	result := make([]entity.ContractPaymentDto, len(paginationResult))
	for i, m := range paginationResult {
		result[i] = new(entity.ContractPaymentDto).FromModel(m)
	}
	info.Data = result
	return info, nil
}

func (r *repository) Update(ctx context.Context, dto *entity.ContractPaymentDto) (*entity.ContractPaymentDto, error) { // Implementation of updating a funder in the database
	err := r.db.Save(dto.ToModel()).Error
	if err != nil {
		return nil, err
	}
	return dto, nil
}

func (r *repository) Delete(ctx context.Context, id string) error {
	err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.ContractPayment{}).Error
	if err != nil {
		return err
	}
	return err
}
