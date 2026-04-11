package model

import (
	"time"

	concern "github.com/raymondsugiarto/funder-api/pkg/model/common"
)

type Contract struct {
	concern.CommonWithIDs
	FunderID           string
	Funder             *Funder
	ContractNumber     int
	ContractCode       string
	DisbursementAt     *time.Time
	Amount             float64
	Duration           int
	DueDate            *time.Time
	DestinationAccount string
	ReturnPercentage   float64
	ReturnAmount       float64
	AttachmentURL      string
	Notes              string
	ContractPayments   []ContractPayment
}
