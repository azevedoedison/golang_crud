package request

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,containsany=!@#?"`
	Name     string `json:"name" binding:"required,max=10"`
	Age      int8   `json:"age" binding:"required"`
}
