package domain

import "errors"

type AppError error

var (
	ErrTimeTrialNotFound    AppError = errors.New("time trial not found")
	ErrBoatNotFound         AppError = errors.New("boat not found")
	ErrOrganizationNotFound AppError = errors.New("organization not found")
	ErrClubNotFound         AppError = errors.New("club not found")
	ErrGroupNotFound        AppError = errors.New("group not found")
	ErrRowerNotFound        AppError = errors.New("rower not found")
	ErrShellNotFound        AppError = errors.New("shell not found")
	ErrRentalNotFound       AppError = errors.New("rental not found")
	ErrRentalRowerNotFound  AppError = errors.New("rental rower not found")
	ErrAuthNotFound         AppError = errors.New("auth not found")
)
