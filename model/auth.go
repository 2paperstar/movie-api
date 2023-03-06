package model

type Credential struct {
	ID       string `json:"id" bson:"_id"`
	Password string `json:"password" bson:"password"`
}

type UserInfo struct {
	Email    string `json:"email" bson:"email"`
	Nickname string `json:"nickname" bson:"nickname"`
	Phone    string `json:"phone" bson:"phone"`
}

type RegisterForm struct {
	Credential `bson:",inline"`
	UserInfo   `bson:",inline"`
}

type RefreshToken struct {
	RefreshToken string `json:"refreshToken"`
}

type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
