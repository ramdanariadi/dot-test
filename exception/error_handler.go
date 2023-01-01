package exception

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			err := recover()

			if errors, ok := err.(validator.ValidationErrors); ok {
				var errMsg []string
				for _, err := range errors {
					errMsg = append(errMsg, err.Field()+" field is "+err.Tag())
				}
				context.JSON(400, gin.H{"message": errMsg})
				return
			}

			if errors, ok := err.(NotFoundError); ok {
				context.JSON(400, gin.H{"message": errors.Message})
				return
			}

			if errors, ok := err.(AuthenticationError); ok {
				context.JSON(403, gin.H{"message": errors.Message})
				return
			}

			if err != nil {
				if errors, ok := err.(error); ok {
					context.JSON(500, gin.H{"message": errors.Error()})
					return
				}
				context.JSON(500, gin.H{"message": err})
				return
			}
		}()

		context.Next()
	}
}
