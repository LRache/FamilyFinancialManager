package response

type UserLogin struct {
	Token string `json:"token" binding:"required"`
}
