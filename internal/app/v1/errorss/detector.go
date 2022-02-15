package errorss

import (
	"github.com/jackc/pgconn"
)

func IsAlreadyExistsErr(err error) bool {
	pgxErr := err.(*pgconn.PgError)
	if pgxErr.Code == "23505" {
		return true
	}
	return false
}