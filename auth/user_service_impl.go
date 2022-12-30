package auth

import (
	_ "embed"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/ramdanariadi/dot-test/exception"
	"github.com/ramdanariadi/dot-test/helpers"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//go:embed JWTSECRET
var jwtSecret []byte

type UserServiceImpl struct {
	DB *gorm.DB
}

func NewUserService(DB *gorm.DB) *UserServiceImpl {
	return &UserServiceImpl{DB: DB}
}

func (u *UserServiceImpl) Login(userDTO UserDTO) Token {
	var user User
	first := u.DB.First(&user, "username = ?", userDTO.Username)
	if first.Error != nil {
		panic(exception.NewAuthenticationError("INVALID_CREDENTIALS"))
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDTO.Password))
	if err != nil {
		panic(exception.NewAuthenticationError("INVALID_CREDENTIALS"))
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"role":   "user",
	})

	signedString, err := token.SignedString(jwtSecret)
	helpers.PanicIfError(err)

	tokens := Token{
		AccessToken:  signedString,
		RefreshToken: "",
	}
	return tokens
}

func (u *UserServiceImpl) UserExist(username string) bool {
	var user User
	first := u.DB.First(&user, "username = ?", username)
	return first.Error == nil
}

func (u *UserServiceImpl) SignUp(user UserDTO) Token {
	if u.UserExist(user.Username) {
		panic(exception.NewAuthenticationError("INVALID_REGISTER"))
	}
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	helpers.PanicIfError(err)
	id, _ := uuid.NewUUID()
	save := u.DB.Save(&User{ID: id.String(), Username: user.Username, Password: string(password)})
	helpers.PanicIfError(save.Error)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": id,
		"role":   "user",
	})

	signedString, err := token.SignedString(jwtSecret)
	helpers.PanicIfError(err)

	tokens := Token{
		AccessToken:  signedString,
		RefreshToken: "",
	}
	return tokens
}
