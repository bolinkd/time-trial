package domain

import "gopkg.in/volatiletech/null.v6"

type InputValidation interface {
	Validate() error
}

func isValidNullInt(x null.Int) bool {
	return x.Valid && x.Int > 0
}
