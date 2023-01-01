package route

import (
	"github.com/gin-gonic/gin"
	"github.com/ramdanariadi/dot-test/auth"
	"github.com/ramdanariadi/dot-test/category"
	"github.com/ramdanariadi/dot-test/exception"
	"github.com/ramdanariadi/dot-test/product"
	"github.com/ramdanariadi/dot-test/transaction"
	"gorm.io/gorm"
)

func RouterSetup(db *gorm.DB) *gin.Engine {
	categoryController := category.NewCategoryController(db)
	productController := product.NewProductControllerImpl(db)
	transactionController := transaction.NewTransactionControllerImpl(db)
	userController := auth.NewUserController(db)

	engine := gin.New()
	engine.Use(exception.ErrorHandler())

	engine.POST("/register", userController.SignUp)
	engine.POST("/login", userController.Login)
	engine.POST("/userExist", userController.UserExist)

	categoryGroup := engine.Group("/category")
	{
		categoryGroup.GET("/", categoryController.FindAll)
		categoryGroup.GET("/:id", categoryController.FindById)
		categoryGroup.Use(auth.SecureRequest()).POST("/", categoryController.Save)
		categoryGroup.Use(auth.SecureRequest()).PUT("/:id", categoryController.Update)
		categoryGroup.Use(auth.SecureRequest()).DELETE("/:id", categoryController.Delete)
	}

	productGroup := engine.Group("/product")
	{
		productGroup.GET("/", productController.FindAll)
		productGroup.GET("/:id", productController.FindById)
		productGroup.Use(auth.SecureRequest()).POST("/", productController.Save)
		productGroup.Use(auth.SecureRequest()).PUT("/:id", productController.Update)
		productGroup.Use(auth.SecureRequest()).DELETE("/:id", productController.Delete)
	}

	transactionGroup := engine.Group("/transaction")
	{
		transactionGroup.Use(auth.SecureRequest()).POST("/", transactionController.Save)
		transactionGroup.Use(auth.SecureRequest()).GET("/:id", transactionController.FindById)
		transactionGroup.Use(auth.SecureRequest()).GET("/", transactionController.Find)
	}
	return engine
}
