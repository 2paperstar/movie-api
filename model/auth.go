package model

type Credential struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

type UserInfo struct {
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
}

type RegisterForm struct {
	Credential
	UserInfo
}

type RefreshToken struct {
	RefreshToken string `json:"refreshToken"`
}

type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
