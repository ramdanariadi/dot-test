package auth

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserControllerImpl struct {
	Service UserService
}

func NewUserController(db *gorm.DB) *UserControllerImpl {
	return &UserControllerImpl{
		Service: NewUserService(db),
	}
}

func (u *UserControllerImpl) Login(ctx *gin.Context) {
	var user UserDTO
	if err := ctx.Bind(&user); err == nil {
		token := u.Service.Login(user)
		ctx.JSON(200, token)
	}
}

func (u *UserControllerImpl) SignUp(ctx *gin.Context) {
	var user UserDTO
	if err := ctx.Bind(&user); err == nil {
		token := u.Service.SignUp(user)
		ctx.JSON(200, token)
	}
}

func (u *UserControllerImpl) UserExist(ctx *gin.Context) {
	var user UserDTO
	if err := ctx.Bind(&user); err == nil {
		ctx.JSON(200, gin.H{
			"exist": u.Service.UserExist(user.Username),
		})
	}
}
