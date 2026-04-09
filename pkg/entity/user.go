package entity

import (
	"github.com/raymondsugiarto/funder-api/pkg/model"
	"github.com/raymondsugiarto/funder-api/shared/security"
)

// to be
type UserDto struct {
	ID              string              `json:"id"`
	UserCredentials []UserCredentialDto `json:"userCredential"`
}

func NewUserDtoFromModel(m *model.User) *UserDto {
	return &UserDto{
		ID:              m.ID,
		UserCredentials: []UserCredentialDto{},
	}
}

func (f *UserDto) FromModel(m *model.User) UserDto {
	return *NewUserDtoFromModel(m)
}

func (f *UserDto) ToModel() *model.User {
	m := &model.User{}
	if len(f.UserCredentials) > 0 {
		m.UserCredentials = make([]model.UserCredential, len(f.UserCredentials))
		for _, uc := range f.UserCredentials {
			m.UserCredentials = append(m.UserCredentials, *uc.ToModel())
		}
	}
	if f.ID != "" {
		m.ID = f.ID
	}
	return m
}

// UserCredentialDto
type UserCredentialDto struct {
	ID             string `json:"id"`
	OrganizationID string `json:"organizationId"`
	Username       string `json:"username"`
	Password       string `json:"-"`
}

func NewUserCredentialDtoFromModel(m *model.UserCredential) *UserCredentialDto {
	return &UserCredentialDto{
		ID:             m.ID,
		OrganizationID: m.OrganizationID,
		Username:       m.Username,
	}
}

func (f *UserCredentialDto) FromModel(m *model.UserCredential) UserCredentialDto {
	return *NewUserCredentialDtoFromModel(m)
}

func (f *UserCredentialDto) ToModel() *model.UserCredential {
	encryptedPass, _ := security.HashPassword(f.Password)
	m := &model.UserCredential{
		OrganizationID: f.OrganizationID,
		Username:       f.Username,
		Password:       encryptedPass,
	}
	if f.ID != "" {
		m.ID = f.ID
	}
	return m
}
