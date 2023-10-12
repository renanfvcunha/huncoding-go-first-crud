package request

type UserRequest struct {
	Email    string `json:"email,omitempty" binding:"required,email"`
	Password string `json:"password,omitempty" binding:"required,min=6,containsany=!@#$%*"`
	Name     string `json:"name,omitempty" binding:"required,min=4,max=100"`
	Age      int8   `json:"age,omitempty" binding:"required,min=1,max=140"`
}
