package cart

import (
	"context"
	"database/sql"
)

type CartRepository interface {
	FindByUserId(context context.Context, tx *sql.Tx, userId string) []CartModel
	Save(context context.Context, tx *sql.Tx, product CartModel) bool
	Delete(context context.Context, tx *sql.Tx, userId string, productId string) bool
}
