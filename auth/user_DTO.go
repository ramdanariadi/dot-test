package auth

type UserDTO struct {
	Username string `verified:"required"`
	Password string `verified:"required"`
}
