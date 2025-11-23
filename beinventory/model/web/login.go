package web

type LoginRequest struct {
	UserOrEmail string `validate:"required,max=100,min=1" json:"useroremail"`
	Password    string `validate:"required,max=100,min=1" json:"password"`
}
