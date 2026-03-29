package model

import concern "github.com/raymondsugiarto/funder-api/pkg/model/common"

// Accounts : table accounts
type Organization struct {
	concern.CommonWithIDs
	Code   string
	Name   string
	Origin string
}
