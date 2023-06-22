package authdto

type AuthRequest struct {
	UserName string `json:"username" form:"username" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Phone    string `json:"phone" form:"password" validate:"required"`
	Address  string `json:"address" form:"address" validate:"required"`
}

type LoginRequest struct {
	UserName string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}
