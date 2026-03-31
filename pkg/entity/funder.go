package entity

import (
	"github.com/raymondsugiarto/funder-api/pkg/model"
	"github.com/raymondsugiarto/funder-api/shared/database/pagination"
)

type FunderRequest struct {
	Name           string `json:"name"`
	PhoneNumber    string `json:"phoneNumber"`
	FunderIDParent string `json:"funderIdParent,omitempty"`
	Password       string `json:"password"`
}

func (r *FunderRequest) ToDto() *FunderDto {
	return &FunderDto{
		Name:           r.Name,
		PhoneNumber:    r.PhoneNumber,
		FunderIDParent: r.FunderIDParent,
		Password:       r.Password,
	}
}

type FunderDto struct {
	ID             string   `json:"id"`
	UserID         string   `json:"userId"`
	User           *UserDto `json:"user,omitempty"`
	Name           string   `json:"name"`
	PhoneNumber    string   `json:"phoneNumber"`
	FunderIDParent string   `json:"funderIdParent,omitempty"`
	Password       string   `json:"-"`
}

func NewFunderDtoFromModel(m *model.Funder) *FunderDto {
	if m == nil {
		return nil
	}

	f := &FunderDto{
		ID:             m.ID,
		UserID:         m.UserID,
		Name:           m.Name,
		PhoneNumber:    m.PhoneNumber,
		FunderIDParent: m.FunderIDParent,
	}
	if m.User != nil {
		f.User = NewUserDtoFromModel(m.User)
	}
	return f
}

func (f *FunderDto) FromModel(m *model.Funder) FunderDto {
	return *NewFunderDtoFromModel(m)
}

func (f *FunderDto) ToModel() *model.Funder {
	m := &model.Funder{
		UserID:         f.UserID,
		Name:           f.Name,
		PhoneNumber:    f.PhoneNumber,
		FunderIDParent: f.FunderIDParent,
	}

	if f.ID != "" {
		m.ID = f.ID
	}

	return m
}

func (f FunderDto) ToUserDto() *UserDto {
	return &UserDto{
		ID: f.UserID,
		UserCredentials: []UserCredentialDto{
			{
				Username: f.PhoneNumber,
				Password: f.Password,
			},
		},
	}
}

type FunderFilterDto struct {
	pagination.GetListRequest
}

func (f FunderFilterDto) AddFilter(filter pagination.FilterItem) {}
