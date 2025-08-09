package registers

type RegisterRequestDto struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Username string `form:"username" json:"username" binding:"required"`
	Name     string `form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,min=8"`
}
