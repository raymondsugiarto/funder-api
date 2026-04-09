package entity

type LoginRequestDto struct {
	Username string `json:"username" bson:"username" validate:"required"`
	Password string `json:"password" bson:"password" validate:"required"`
}

type LoginDto struct {
	Token   string `json:"token" bson:"token"`
	Expired string `json:"expired" bson:"exp"`
}

type UserSessionDto struct {
	ID         string `json:"id"`  // user credential id
	UserID     string `json:"uid"` // user id
	CustomerID string `json:"cid"` // user id
	AccountID  string `json:"aid"` // user id
}
