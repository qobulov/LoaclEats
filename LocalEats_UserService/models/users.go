package models

type UserInfo struct {
	Id           string
	Username     string
	Password     string
	Phone_number string
}

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
