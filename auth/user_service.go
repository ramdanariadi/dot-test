package auth

type UserService interface {
	Login(userDTO UserDTO) Token
	SignUp(userDTO UserDTO) Token
	UserExist(username string) bool
}
