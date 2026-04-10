package model

import concern "github.com/raymondsugiarto/funder-api/pkg/model/common"

type Funder struct {
	concern.CommonWithIDs
	UserID         string
	User           *User `gorm:"foreignKey:UserID;references:ID"`
	Name           string
	PhoneNumber    string
	FunderIDParent string
}
