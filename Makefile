save:
	rm -rf Godeps/ vendor/
	godep save ./...

sqlboiler:
	sqlboiler -b goose_db_version --wipe --no-tests postgres

test:
	go test ./...

# Creates mocks for testing
mocks:
	cd ./socket &&\
	mockery -name=ClientInterface

	cd ./onspot &&\
	mockery -name=ClientInterface

	cd ./paytrace &&\
    	mockery -name=ClientInterface

	cd ./appnexus &&\
	mockery -name=ClientInterface

	cd ./aws &&\
	mockery -name=ClientInterface

	cd ./db &&\
    	mockery -all

# Get or updates the dependencies for this project
deps:
	go get -u github.com/gin-gonic/gin
	go get -u github.com/volatiletech/sqlboiler

# Cleans our project: deletes binaries and coverage report
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
	if [ -f coverage.html ] ; then rm coverage.html ; fi

.PHONY: clean test
