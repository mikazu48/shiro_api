package request

type LoginRequest struct {
	Username string `validate:"required,max=50,min=5" json:"username"`
	Password string `validate:"required,max=50,min=5" json:"password"`
}
