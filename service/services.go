package service

type Services struct{}

type ServicesInterface interface {
	TimeTrialServiceInterface
	BoatServiceInterface
	OrganizationServiceInterface
	ClubServiceInterface
	GroupServiceInterface
	RowerServiceInterface
	ShellServiceInterface
	RentalServiceInterface
	RentalRowerServiceInterface
}

// Create creates the new db server connection
func Create() Services {
	return Services{}
}

// Close the database connection
func (conn *Services) Close() {}
