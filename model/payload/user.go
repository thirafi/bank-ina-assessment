package payload

type (
	CreateUserRequest struct {
		Name     string `json:"name" form:"name" validate:"required,max=20"`
		Email    string `json:"email" form:"email" validate:"required,email"`
		Password string `json:"password" form:"password" validate:"required,min=5"`
	}

	CreateUserResponse struct {
		UserID uint   `json:"user_id"`
		Token  string `json:"token"`
	}

	LoginUserRequest struct {
		Email    string `json:"email" form:"email" validate:"required,email"`
		Password string `json:"password" form:"password" validate:"required,min=5"`
	}

	LoginUserResponse struct {
		Email       string `json:"email"`
		Name        string `json:"name"`
		AccessToken string `json:"access_token"`
	}
)
