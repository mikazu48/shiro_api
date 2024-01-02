package request

type UserRequest struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	FullName string `json:"full_name" form:"full_name"`
}
