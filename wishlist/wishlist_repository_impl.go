package wishlist

import (
	"database/sql"
)

type WishlistRepositoryImpl struct {
	DB *sql.DB
}
