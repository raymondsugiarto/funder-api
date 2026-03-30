package entity

import (
	"github.com/raymondsugiarto/funder-api/pkg/model"
	"github.com/raymondsugiarto/funder-api/shared/pagination"
)

type FunderRequest struct {
	Name           string `json:"name"`
	PhoneNumber    string `json:"phoneNumber"`
	FunderIDParent string `json:"funderIdParent,omitempty"`
}

func (r *FunderRequest) ToDto() *FunderDto {
	return &FunderDto{
		Name:           r.Name,
		PhoneNumber:    r.PhoneNumber,
		FunderIDParent: r.FunderIDParent,
	}
}

type FunderDto struct {
	ID             string
	UserID         string
	User           *UserDto
	Name           string
	PhoneNumber    string
	FunderIDParent string
}

func NewFunderDtoFromModel(funder *model.Funder) *FunderDto {
	if funder == nil {
		return nil
	}

	return &FunderDto{
		ID:             funder.ID,
		UserID:         funder.UserID,
		User:           NewUserDtoFromModel(funder.User),
		Name:           funder.Name,
		PhoneNumber:    funder.PhoneNumber,
		FunderIDParent: funder.FunderIDParent,
	}
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

type FunderFilterDto struct {
	pagination.GetListRequest
}

func (f FunderFilterDto) AddFilter(filter pagination.FilterItem) {}
