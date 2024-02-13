package dto

type CreateCustomerRequestDTO struct {
	Username string `validate:"required,min=1,max=100" json:"username"`
	Password string `validate:"required,min=1,max=100" json:"password"`
}

type CreateCustomerResponseDTO struct {
	Username string `json:"username"`
}

type LoginCustomerRequestDTO struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}

type LoginCustomerResponseDTO struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
