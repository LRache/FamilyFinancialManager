package response

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	ID       int    `json:"id" binding:"required"`
	Token    string `json:"token" binding:"required"`
}
