package entity

import (
	"mime/multipart"
	"time"

	"github.com/raymondsugiarto/funder-api/pkg/model"
	"github.com/raymondsugiarto/funder-api/shared/database/pagination"
)

type ContractPaymentRequest struct {
	ContractID     string                `json:"contractId"`
	PaymentAt      *time.Time            `json:"paymentAt"`
	PaymentAmount  float64               `json:"paymentAmount"`
	AttachmentFile *multipart.FileHeader `json:"attachmentFile"`
	Notes          string                `json:"notes"`
}

func (r *ContractPaymentRequest) ToDto(attachmentUrl string) *ContractPaymentDto {
	return &ContractPaymentDto{
		ContractID:    r.ContractID,
		PaymentAt:     r.PaymentAt,
		PaymentAmount: r.PaymentAmount,
		AttachmentURL: attachmentUrl,
		Notes:         r.Notes,
	}
}

type ContractPaymentDto struct {
	ID            string       `json:"id"`
	ContractID    string       `json:"contractId"`
	Contract      *ContractDto `json:"contract,omitempty"`
	PaymentAt     *time.Time   `json:"paymentAt"`
	PaymentAmount float64      `json:"paymentAmount"`
	AttachmentURL string       `json:"attachmentUrl"`
	Notes         string       `json:"notes"`
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
	FunderID string
}

func (f *ContractPaymentFilterDto) GenerateFilter() {
	if f.FunderID != "" {
		f.AddFilter(pagination.FilterItem{
			Field: "funder_id",
			Op:    "eq",
			Val:   f.FunderID,
		})
	}
}
