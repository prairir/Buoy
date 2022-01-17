# vendor dependencies
deps:
	FROM golang:1.17-bullseye
	WORKDIR /app
	COPY go.mod go.sum .
	RUN go mod download

# build the binary
build:
	# build off of deps stage
	FROM +deps
	# copy main file
	COPY main.go .
	# copy cmd and whatever else in dir mode
	# this is like `cp -r`
	COPY --dir cmd/ pkg/ ./
	# build to file `imacry`
	RUN go build -race -o buoy main.go
	# save file as artifact
	SAVE ARTIFACT buoy AS LOCAL buoy

# run the tests
test:
	FROM +build

	RUN --privileged go test ./...
