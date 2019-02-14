package domain

import "github.com/volatiletech/null"

type InputValidation interface {
	Validate() error
}

func isValidNullInt(x null.Int) bool {
	return x.Valid && x.Int > 0
}
