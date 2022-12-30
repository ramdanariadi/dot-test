package auth

import (
	"github.com/gin-gonic/gin"
)

type UserController interface {
	Login(ctx *gin.Context)
	SignUp(ctx *gin.Context)
	UserExist(ctx *gin.Context)
}
