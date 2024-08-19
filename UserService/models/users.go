package models

type RefreshToken struct {
	UserId    string
	Email     string
	Token     string
	CreatedAt int64
	ExpiresAt int64
}

type Error struct {
	Error string `json:"error"`
}
