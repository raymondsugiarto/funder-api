package usercredential

import (
	"context"
	"errors"

	"github.com/raymondsugiarto/funder-api/pkg/entity"
	"github.com/raymondsugiarto/funder-api/pkg/model"
	"github.com/raymondsugiarto/funder-api/shared/database/pagination"

	"gorm.io/gorm"
)

type Repository interface {
	FindByUsername(ctx context.Context, dto *entity.UserCredentialDto) (*entity.UserCredentialDto, error)
	FindByEmail(ctx context.Context, dto *entity.UserCredentialDto) (*entity.UserCredentialDto, error)
	GetUserCredentialByUsername(sctx context.Context, organizationID string, username string) (*model.UserCredential, error)
	ChangePassword(ctx context.Context, userId, password string) error

	Create(ctx context.Context, dto *entity.UserCredentialDto) (*entity.UserCredentialDto, error)
	FindByID(ctx context.Context, id string) (*entity.UserCredentialDto, error)
	FindAll(ctx context.Context, req pagination.PaginationRequestDto) (*pagination.ResultPagination[entity.UserCredentialDto], error)
	Update(ctx context.Context, dto *entity.UserCredentialDto) (*entity.UserCredentialDto, error)
	Delete(ctx context.Context, id string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, dto *entity.UserCredentialDto) (*entity.UserCredentialDto, error) {
	m := dto.ToModel()
	err := r.db.Create(m).Error
	if err != nil {
		return nil, err
	}
	dto.ID = m.ID
	return dto, nil
}

func (r *repository) FindByID(ctx context.Context, id string) (*entity.UserCredentialDto, error) {
	var m entity.UserCredentialDto
	err := r.db.WithContext(ctx).First(&m, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *repository) FindAll(ctx context.Context, req pagination.PaginationRequestDto) (*pagination.ResultPagination[entity.UserCredentialDto], error) {
	info, paginationResult, err := pagination.NewTable[*entity.FunderFilterDto, *model.UserCredential, entity.UserCredentialDto]().
		Paginate(ctx, func(req *entity.FunderFilterDto) *gorm.DB {
			query := r.db.WithContext(ctx).Model(&model.UserCredential{})
			return query
		}, &pagination.TableRequest[*entity.FunderFilterDto]{})
	if err != nil {
		return nil, err
	}
	result := make([]entity.UserCredentialDto, len(paginationResult))
	for i, m := range paginationResult {
		result[i] = new(entity.UserCredentialDto).FromModel(m)
	}
	info.Data = result
	return info, nil
}

func (r *repository) Update(ctx context.Context, dto *entity.UserCredentialDto) (*entity.UserCredentialDto, error) { // Implementation of updating a funder in the database
	return nil, nil
}

func (r *repository) Delete(ctx context.Context, id string) error {
	// Implementation of deleting a funder from the database
	return nil
}

// ChangePassword is a function to change user password
func (r *repository) ChangePassword(ctx context.Context, userId, password string) error {
	var userCredential model.UserCredential
	if err := r.db.WithContext(ctx).Model(&userCredential).
		Where("user_id = ?", userId).
		Update("password", password).Error; err != nil {
		return err
	}
	return nil
}

// FindByUsername is a function to find user credential by username
func (r *repository) FindByUsername(ctx context.Context, userCredential *entity.UserCredentialDto) (*entity.UserCredentialDto, error) {
	var userCredentialModel model.UserCredential
	if err := r.db.Joins("User").
		Where("user_credential.username = ? AND user_credential.organization_id = ?", userCredential.Username, userCredential.OrganizationID).
		First(&userCredentialModel).Error; err != nil {
		return nil, err
	}
	userCredential.ID = userCredentialModel.ID
	return userCredential, nil
}

// FindByUsername is a function to find user credential by username
func (r *repository) FindByEmail(ctx context.Context, userCredential *entity.UserCredentialDto) (*entity.UserCredentialDto, error) {
	var userCredentialModel model.UserCredential
	if err := r.db.Joins("User").
		Where("user_credential.username = ? AND user_credential.organization_id = ?", userCredential.Username, userCredential.OrganizationID).
		First(&userCredentialModel).Error; err != nil {
		return nil, err
	}
	userCredential.ID = userCredentialModel.ID
	return userCredential, nil
}

func (r *repository) GetUserCredentialByUsername(ctx context.Context, organizationID, u string) (*model.UserCredential, error) {
	var userCredential model.UserCredential
	if err := r.db.Joins("User").
		Where(&model.UserCredential{OrganizationID: organizationID, Username: u}).
		Find(&userCredential).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("userNotFound")
		}
		return nil, err
	}
	return &userCredential, nil
}
