package model

import concern "github.com/raymondsugiarto/funder-api/pkg/model/common"

type Funder struct {
	concern.CommonWithIDs
	UserID         string `gorm:"type:varchar(255);not null"`
	FunderIDParent string `gorm:"type:varchar(255)"`
}
