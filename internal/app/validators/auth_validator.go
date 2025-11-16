package validators

type RegisterValidator struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginValidator struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RefreshValidator struct {
	Email        string `json:"email" validate:"required,email"`
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type LogoutValidator struct {
	Email string `json:"email" validate:"required,email"`
}
