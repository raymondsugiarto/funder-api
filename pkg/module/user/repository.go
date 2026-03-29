package user

import (
	"context"

	"github.com/raymondsugiarto/funder-api/pkg/entity"
	"github.com/raymondsugiarto/funder-api/pkg/model"
	"gorm.io/gorm"
)

type Repository interface {
	// Create(ctx context.Context, dto *entity.UserDto) (*entity.UserDto, error)
	// FindByReferralCode(string) (*entity.CreateUser, error)
	FindByID(ctx context.Context, id string) (*entity.UserDto, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// FindByUserID is a function to find user by user id
func (r *repository) FindByID(ctx context.Context, id string) (*entity.UserDto, error) {
	var user *model.User
	if err := r.db.WithContext(ctx).Model(&model.User{}).
		Where("id = ?", id).
		First(&user).Error; err != nil {
		return nil, err
	}

	return entity.NewUserDtoFromModel(user), nil
}

// // FindByReferralCode is a function to find user by referral code
// func (r *repository) FindByReferralCode(referralCode string) (*entity.CreateUser, error) {
// 	var user model.User
// 	if err := r.db.Joins("Customer").
// 		Where("customer.referral_code = ?", referralCode).
// 		First(&user).Error; err != nil {
// 		return nil, err
// 	}
// 	return &entity.CreateUser{}, nil
// }

// // CreateUser is a function to create user
// func (r *repository) CreateUser(createUser *entity.CreateUser) (*entity.CreateUser, error) {
// 	user := new(model.User)
// 	user.OrganizationID = createUser.OrganizationData.ID

// 	password, _ := utils.HashPassword(createUser.Password)
// 	user.UserCredential = []model.UserCredential{
// 		{
// 			OrganizationID: createUser.OrganizationData.ID,
// 			Username:       createUser.Username,
// 			Password:       password,
// 		},
// 	}

// 	if err := r.db.Create(user).Error; err != nil {
// 		return nil, err
// 	}
// 	createUser.UserID = user.ID
// 	return createUser, nil
// }
