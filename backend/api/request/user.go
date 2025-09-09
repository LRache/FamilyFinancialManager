package request

type UserRegister struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Email    string `form:"email" binding:"required"`
}

type UserLogin struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
