package model

import (
	concern "github.com/raymondsugiarto/funder-api/pkg/model/common"
)

type UserType string

const (
	ADMIN  UserType = "ADMIN"
	FUNDER UserType = "FUNDER"
)

type User struct {
	concern.CommonWithIDs
	UserType        UserType
	UserCredentials []UserCredential
}
