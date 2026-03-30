package entity

import (
	"github.com/raymondsugiarto/funder-api/pkg/model"
)

type CustomerAccountListItem struct {
	ID                 string `json:"id" bson:"id"`
	AccountCode        string `json:"accountCode" bson:"accountCode"`
	UserID             string `json:"userId" bson:"userId"`
	CreatedAt          string `json:"createdAt" bson:"createdAt"`
	CustomerFollowerID string `json:"customerFollowerId" bson:"customerFollowerId"`
}

type MyAccountProfile struct {
	ID             string `json:"id" bson:"id"`
	AccountCode    string `json:"accountCode" bson:"accountCode"`
	Email          string `json:"email" bson:"email"`
	PhoneNumber    string `json:"phoneNumber" bson:"phoneNumber"`
	FirstName      string `json:"firstName" bson:"firstName"`
	LastName       string `json:"lastName" bson:"lastName"`
	UserID         string `json:"userId" bson:"userId"`
	CreatedAt      string `json:"createdAt" bson:"createdAt"`
	FollowerCount  int    `json:"followerCount" bson:"followerCount"`
	FollowingCount int    `json:"followingCount" bson:"followingCount"`
}

type CustomerAccount struct {
	ID          string `json:"id" bson:"id"`
	AccountCode string `json:"accountCode" bson:"accountCode"`
	UserID      string `json:"userId" bson:"userId"`
	CreatedAt   string `json:"createdAt" bson:"createdAt"`
}

type CustomerFollower struct {
	ID                        string `json:"id" bson:"id"`
	CustomerAccountID         string `json:"customerAccountId" bson:"customerAccountId"`
	CustomerAccountIDFollower string `json:"customerAccountIdFollower" bson:"customerAccountIdFollower"`
}

type UserCredentialDto struct {
	ID             string `json:"id"`
	OrganizationID string `json:"organizationId"`
	Username       string `json:"username"`
	Email          string `json:"email" `
}


type UserDto struct {
}

func NewUserDtoFromModel(m *model.User) *UserDto {
	return &UserDto{}
}

