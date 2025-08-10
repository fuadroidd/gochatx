package dtos

type UserRequestDTO struct {
	Username    string `form:"username"`
	Displayname string `form:"displayname"`
	Email       string `form:"email"`
	Password    string `form:"password"`
}

type UserResponseDTO struct {
	ID          uint   `json:"id"`
	Username    string `json:"user_name"`
	Displayname string `json:"display_name"`
	Email       string `json:"email"`
}

type UserLoginRequestDTO struct {
	Username string `form:"username"`
	Password string `fomr:"password"`
}
