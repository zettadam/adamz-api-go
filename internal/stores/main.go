package stores

import (
	"errors"

	"github.com/jackc/pgx/v5"
)

func IsNotFound(err error) bool {
	return errors.Is(err, pgx.ErrNoRows)
}
