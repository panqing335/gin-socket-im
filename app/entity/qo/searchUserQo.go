package qo

type SearchUserQo struct {
	Nickname string `json:"nickname" form:"nickname"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
}
