GOCC=go
# To Compile the linux version using docker simply invoke the makefile like this:
#
# make GOCC="docker run --rm -t -v ${GOPATH}:/go hbouvier/go-lang:1.5"
PROJECTNAME=httpclient

all: get-deps fmt macos linux arm test coverage

clean:
	rm -rf coverage.out \
	       ${GOPATH}/pkg/{linux_amd64,darwin_amd64,linux_arm}/github.com/hbouvier/${PROJECTNAME}

build: fmt test
	${GOCC} install github.com/hbouvier/${PROJECTNAME}

fmt:
	${GOCC} fmt github.com/hbouvier/${PROJECTNAME}

test:
	${GOCC} test -v -cpu 4 -count 1 -coverprofile=coverage.out github.com/hbouvier/${PROJECTNAME}

coverage:
	${GOCC} tool cover -html=coverage.out

get-deps:
	# ${GOCC} get ...

linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 ${GOCC} install github.com/hbouvier/${PROJECTNAME}

macos:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 ${GOCC} install github.com/hbouvier/${PROJECTNAME}

arm:
	GOOS=linux GOARCH=arm CGO_ENABLED=0 ${GOCC} install github.com/hbouvier/${PROJECTNAME}
