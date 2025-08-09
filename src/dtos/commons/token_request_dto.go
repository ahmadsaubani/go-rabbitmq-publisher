package commons

type TokenRequestDto struct {
	Token string `form:"token" json:"token" binding:"required"`
}
