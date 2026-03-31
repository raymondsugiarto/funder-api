package entity

import (
	"time"

	"github.com/raymondsugiarto/funder-api/pkg/model"
	"github.com/raymondsugiarto/funder-api/shared/database/pagination"
)

type ContractPaymentRequest struct {
	ContractID    string     `json:"contractId"`
	PaymentAt     *time.Time `json:"paymentAt"`
	PaymentAmount float64    `json:"paymentAmount"`
	AttachmentURL string     `json:"attachmentUrl"`
	Notes         string     `json:"notes"`
}

func (r *ContractPaymentRequest) ToDto() *ContractPaymentDto {
	return &ContractPaymentDto{
		ContractID:    r.ContractID,
		PaymentAt:     r.PaymentAt,
		PaymentAmount: r.PaymentAmount,
		AttachmentURL: r.AttachmentURL,
		Notes:         r.Notes,
	}
}

type ContractPaymentDto struct {
	ID            string
	ContractID    string
	Contract      *ContractDto
	PaymentAt     *time.Time
	PaymentAmount float64
	AttachmentURL string
	Notes         string
}

func NewContractPaymentDtoFromModel(ContractPayment *model.ContractPayment) *ContractPaymentDto {
	if ContractPayment == nil {
		return nil
	}

	return &ContractPaymentDto{
		ID:            ContractPayment.ID,
		ContractID:    ContractPayment.ContractID,
		Contract:      NewContractDtoFromModel(ContractPayment.Contract),
		PaymentAt:     ContractPayment.PaymentAt,
		PaymentAmount: ContractPayment.PaymentAmount,
		AttachmentURL: ContractPayment.AttachmentURL,
		Notes:         ContractPayment.Notes,
	}
}

func (f *ContractPaymentDto) FromModel(m *model.ContractPayment) ContractPaymentDto {
	return *NewContractPaymentDtoFromModel(m)
}

func (f *ContractPaymentDto) ToModel() *model.ContractPayment {
	m := &model.ContractPayment{
		ContractID:    f.ContractID,
		PaymentAt:     f.PaymentAt,
		PaymentAmount: f.PaymentAmount,
		AttachmentURL: f.AttachmentURL,
		Notes:         f.Notes,
	}

	if f.ID != "" {
		m.ID = f.ID
	}

	return m
}

type ContractPaymentFilterDto struct {
	pagination.GetListRequest
}
