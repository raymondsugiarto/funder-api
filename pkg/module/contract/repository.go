package funder

import (
	"context"

	"github.com/raymondsugiarto/funder-api/pkg/entity"
	"github.com/raymondsugiarto/funder-api/pkg/model"
	"github.com/raymondsugiarto/funder-api/shared/database/pagination"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, dto *entity.ContractDto) (*entity.ContractDto, error)
	FindByID(ctx context.Context, id string) (*entity.ContractDto, error)
	FindAll(ctx context.Context, req pagination.PaginationRequestDto) (*pagination.ResultPagination[entity.ContractDto], error)
	Update(ctx context.Context, dto *entity.ContractDto) (*entity.ContractDto, error)
	Delete(ctx context.Context, id string) error
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
	var m entity.ContractDto
	err := r.db.WithContext(ctx).First(&m, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *repository) FindAll(ctx context.Context, req pagination.PaginationRequestDto) (*pagination.ResultPagination[entity.ContractDto], error) {
	info, paginationResult, err := pagination.NewTable[*entity.ContractFilterDto, *model.Contract, entity.ContractDto]().
		Paginate(ctx, func(req *entity.ContractFilterDto) *gorm.DB {
			query := r.db.WithContext(ctx).Model(&model.Contract{})
			return query
		}, &pagination.TableRequest[*entity.ContractFilterDto]{})
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

func (r *repository) Update(ctx context.Context, dto *entity.ContractDto) (*entity.ContractDto, error) { // Implementation of updating a funder in the database
	return nil, nil
}

func (r *repository) Delete(ctx context.Context, id string) error {
	// Implementation of deleting a funder from the database
	return nil
}
