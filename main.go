package main

import (
	"github.com/ramdanariadi/dot-test/auth"
	"github.com/ramdanariadi/dot-test/category"
	"github.com/ramdanariadi/dot-test/helpers"
	"github.com/ramdanariadi/dot-test/product"
	"github.com/ramdanariadi/dot-test/route"
	"github.com/ramdanariadi/dot-test/setup"
	"github.com/ramdanariadi/dot-test/transaction"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	connection, err := setup.NewDbConnection()
	helpers.PanicIfError(err)

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: connection}))
	helpers.PanicIfError(err)

	err = db.AutoMigrate(&category.Category{}, &product.Product{}, &auth.User{}, &transaction.Transaction{}, &transaction.DetailTransaction{})
	helpers.LogIfError(err)

	engine := route.RouterSetup(db)
	err = engine.Run()
	helpers.LogIfError(err)
}
