all: test build

build:
	@cd cmd/commenttags && go build
	@./cmd/commenttags/commenttags -v
	@echo "BUILD & EXECUTABLE"

test:
	@go test -v
