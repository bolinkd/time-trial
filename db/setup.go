package db

var (
	connection DatabaseInterface       = Connection{}
	_          BoatDBInterface         = connection
	_          ClubDBInterface         = connection
	_          GroupDBInterface        = connection
	_          OrganizationDBInterface = connection
	_          RentalRowerDBInterface  = connection
	_          RentalDBInterface       = connection
	_          RowerDBInterface        = connection
	_          ShellDBInterface        = connection
	_          TimeTrialDBInterface    = connection
	_          AuthDBInterface         = connection
)
