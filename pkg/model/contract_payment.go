package model

import (
	"time"

	concern "github.com/raymondsugiarto/funder-api/pkg/model/common"
)

type ContractPayment struct {
	concern.CommonWithIDs
	ContractID    string
	Contract      *Contract
	PaymentAt     *time.Time
	PaymentAmount float64
	AttachmentURL string
	Notes         string
}
