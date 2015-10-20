all: test build

build:
	@cd cmd/commenttags && go build
	@./cmd/commenttags/commenttags -v
	@echo "BUILD & EXECUTABLE"

test:
	@echo "### TESTING PACKAGE ##################################################"
	@go test -v -cover
	@echo "### TESTING CMD ######################################################"
	@cd cmd/commenttags && go test -v -cover
