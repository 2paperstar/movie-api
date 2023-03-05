package model

type Credential struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

type RefreshToken struct {
	RefreshToken string `json:"refreshToken"`
}

type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
