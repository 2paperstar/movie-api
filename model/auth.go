package model

type Credential struct {
	ID       string `json:"id" bson:"id"`
	Password string `json:"password" bson:"password"`
}

type UserInfoForm struct {
	Email    string `json:"email" bson:"email"`
	Nickname string `json:"nickname" bson:"nickname"`
	Phone    string `json:"phone" bson:"phone"`
}

type UserInfo struct {
	UserInfoForm `bson:",inline"`
	ID           string `json:"id" bson:"id"`
	UID          string `json:"uid" bson:"_id"`
}

type RegisterForm struct {
	Credential   `bson:",inline"`
	UserInfoForm `bson:",inline"`
}

type RefreshToken struct {
	RefreshToken string `json:"refreshToken"`
}

type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
