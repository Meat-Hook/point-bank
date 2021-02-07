package repo

import (
	"database/sql"
	"errors"

	"github.com/Meat-Hook/point-bank/internal/modules/user/internal/app"
)

func convertErr(err error) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return app.ErrNotFound
	default:
		return err
	}
}
