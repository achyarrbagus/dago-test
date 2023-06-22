package authdto

type LoginResponse struct {
	Name   string `json:"username" form:"username"`
	Token  string `json:"token"`
	UserID int    `json:"userId"`
}
