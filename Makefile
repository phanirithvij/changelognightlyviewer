all: build

build:
	go generate -v ./...; go build -o changelognightlyviewer.out .
