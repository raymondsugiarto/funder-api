package entity

import (
	"mime/multipart"
	"time"

	"github.com/raymondsugiarto/funder-api/pkg/model"
	"github.com/raymondsugiarto/funder-api/shared/database/pagination"
)

type ContractRequest struct {
	ContractCode     string                `json:"contractCode"`
	FunderID         string                `json:"funderId"`
	DisbursementAt   *time.Time            `json:"disbursementAt"`
	Amount           float64               `json:"amount"`
	Duration         int                   `json:"duration"`
	ReturnPercentage float64               `json:"returnPercentage"`
	ReturnAmount     float64               `json:"returnAmount"`
	AttachmentFile   *multipart.FileHeader `json:"attachmentFile"`
	Notes            string                `json:"notes"`
}

func (r *ContractRequest) ToDto(attachmentUrl string) *ContractDto {
	dueDate := r.DisbursementAt.AddDate(0, r.Duration, 0)
	return &ContractDto{
		ContractCode:     r.ContractCode,
		FunderID:         r.FunderID,
		DisbursementAt:   r.DisbursementAt,
		Amount:           r.Amount,
		Duration:         r.Duration,
		DueDate:          &dueDate,
		ReturnPercentage: r.ReturnPercentage,
		ReturnAmount:     r.ReturnAmount,
		AttachmentURL:    attachmentUrl,
		Notes:            r.Notes,
	}
}

type ContractDto struct {
	ID               string               `json:"id"`
	FunderID         string               `json:"funderId"`
	Funder           *FunderDto           `json:"funder,omitempty"`
	ContractNumber   int                  `json:"contractNumber"`
	ContractCode     string               `json:"contractCode"`
	DisbursementAt   *time.Time           `json:"disbursementAt"`
	Amount           float64              `json:"amount"`
	Duration         int                  `json:"duration"`
	DueDate          *time.Time           `json:"dueDate"`
	ReturnPercentage float64              `json:"returnPercentage"`
	ReturnAmount     float64              `json:"returnAmount"`
	AttachmentURL    string               `json:"attachmentUrl"`
	Notes            string               `json:"notes"`
	ContractPayments []ContractPaymentDto `json:"contractPayments,omitempty"`
}

func NewContractDto() *ContractDto {
	return &ContractDto{}
}

func NewContractDtoFromModel(m *model.Contract) *ContractDto {
	if m == nil {
		return nil
	}

	dto := &ContractDto{
		ID:               m.ID,
		FunderID:         m.FunderID,
		ContractNumber:   m.ContractNumber,
		ContractCode:     m.ContractCode,
		DisbursementAt:   m.DisbursementAt,
		Amount:           m.Amount,
		Duration:         m.Duration,
		DueDate:          m.DueDate,
		ReturnPercentage: m.ReturnPercentage,
		ReturnAmount:     m.ReturnAmount,
		AttachmentURL:    m.AttachmentURL,
		Notes:            m.Notes,
	}

	if m.ContractPayments != nil {
		dto.ContractPayments = make([]ContractPaymentDto, len(m.ContractPayments))
		for i, payment := range m.ContractPayments {
			dto.ContractPayments[i] = *NewContractPaymentDtoFromModel(&payment)
		}
	}

	return dto
}

func (f *ContractDto) FromModel(m *model.Contract) ContractDto {
	return *NewContractDtoFromModel(m)
}

func (f *ContractDto) ToModel() *model.Contract {
	m := &model.Contract{
		FunderID:         f.FunderID,
		ContractNumber:   f.ContractNumber,
		ContractCode:     f.ContractCode,
		DisbursementAt:   f.DisbursementAt,
		Amount:           f.Amount,
		Duration:         f.Duration,
		DueDate:          f.DueDate,
		ReturnPercentage: f.ReturnPercentage,
		ReturnAmount:     f.ReturnAmount,
		AttachmentURL:    f.AttachmentURL,
		Notes:            f.Notes,
	}

	if f.ID != "" {
		m.ID = f.ID
	}

	return m
}

type ContractFilterDto struct {
	pagination.GetListRequest
	FunderID string `query:"funderId"`
}

func (f *ContractFilterDto) GenerateFilter() {
	if f.FunderID != "" {
		f.AddFilter(pagination.FilterItem{
			Field: "funder_id",
			Op:    "eq",
			Val:   f.FunderID,
		})
	}
}
