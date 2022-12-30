package wishlist

import (
	"context"
	"database/sql"
)

type WishlistRepository interface {
	FindByUserId(context context.Context, tx *sql.Tx, userId string)
	FindByUserAndProductId(context context.Context, tx *sql.Tx, userId string, productId string)
	Save(context context.Context, tx *sql.Tx, product WishlistModel) bool
	Delete(context context.Context, tx *sql.Tx, userId string, productId string) bool
}
