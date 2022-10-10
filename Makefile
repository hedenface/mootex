.DEFAULT_GOAL := build

build:
	GOOS=linux go build -o mootex ./src/
