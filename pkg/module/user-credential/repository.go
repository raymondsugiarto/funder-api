package usercredential

import (
	"context"
	"errors"

	"github.com/raymondsugiarto/funder-api/pkg/entity"
	"github.com/raymondsugiarto/funder-api/pkg/model"

	"gorm.io/gorm"
)

type Repository interface {
	FindByUsername(ctx context.Context, dto *entity.UserCredentialDto) (*entity.UserCredentialDto, error)
	FindByEmail(ctx context.Context, dto *entity.UserCredentialDto) (*entity.UserCredentialDto, error)
	GetUserCredentialByUsername(sctx context.Context, organizationID string, username string) (*model.UserCredential, error)
	ChangePassword(ctx context.Context, userId, password string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
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
		Where("user_credential.username = ? AND user_credential.organization_id = ?", userCredential.Email, userCredential.OrganizationID).
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
