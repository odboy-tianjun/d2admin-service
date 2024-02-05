package dto

type KickOutDTO struct {
	Username string `form:"username" json:"username" uri:"username" xml:"username" binding:"required"`
}
