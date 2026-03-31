package funder

import (
	"context"

	"github.com/raymondsugiarto/funder-api/pkg/entity"
	"github.com/raymondsugiarto/funder-api/pkg/model"
	"github.com/raymondsugiarto/funder-api/shared/database/pagination"
	"github.com/raymondsugiarto/funder-api/shared/database/transaction"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, dto *entity.FunderDto) (*entity.FunderDto, error)
	FindByID(ctx context.Context, id string) (*entity.FunderDto, error)
	FindAll(ctx context.Context, req pagination.PaginationRequestDto) (*pagination.ResultPagination[entity.FunderDto], error)
	Update(ctx context.Context, dto *entity.FunderDto) (*entity.FunderDto, error)
	Delete(ctx context.Context, id string) error
}

type repository struct {
	transaction.AppRepository
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, dto *entity.FunderDto) (*entity.FunderDto, error) {
	m := dto.ToModel()
	err := r.db.Create(m).Error
	if err != nil {
		return nil, err
	}
	dto.ID = m.ID
	return dto, nil
}

func (r *repository) FindByID(ctx context.Context, id string) (*entity.FunderDto, error) {
	var m entity.FunderDto
	err := r.db.WithContext(ctx).First(&m, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *repository) FindAll(ctx context.Context, req pagination.PaginationRequestDto) (*pagination.ResultPagination[entity.FunderDto], error) {
	info, paginationResult, err := pagination.NewTable[*entity.FunderFilterDto, *model.Funder, entity.FunderDto]().
		Paginate(ctx, func(req *entity.FunderFilterDto) *gorm.DB {
			query := r.db.WithContext(ctx).Model(&model.Funder{})
			return query
		}, &pagination.TableRequest[*entity.FunderFilterDto]{Request: req.(*entity.FunderFilterDto)})
	if err != nil {
		return nil, err
	}
	result := make([]entity.FunderDto, len(paginationResult))
	for i, m := range paginationResult {
		result[i] = new(entity.FunderDto).FromModel(m)
	}
	info.Data = result
	return info, nil
}

func (r *repository) Update(ctx context.Context, dto *entity.FunderDto) (*entity.FunderDto, error) { // Implementation of updating a funder in the database
	return nil, nil
}

func (r *repository) Delete(ctx context.Context, id string) error {
	// Implementation of deleting a funder from the database
	return nil
}
