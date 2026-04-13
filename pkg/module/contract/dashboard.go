package contract

import (
	"context"

	"github.com/raymondsugiarto/funder-api/pkg/entity"
)

type DashboardService interface {
	ViewDashboard(c context.Context, funderID string) (*entity.DashboardDto, error)
}

func (s *service) ViewDashboard(c context.Context, funderID string) (*entity.DashboardDto, error) {
	return s.repo.ViewDashboard(c, funderID)
}
