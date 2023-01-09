package dto

type EditUserDto struct {
	ID          int64  `json:"id"`
	Nickname    string `json:"nickname"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
}
