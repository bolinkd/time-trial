package domain

import "errors"

type TraxError error

var (
	ErrNoFeatures TraxError = errors.New("no features in featureCollection")
)
