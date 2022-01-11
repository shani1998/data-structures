GO_PACKAGES		:= $(shell go list ./...)

# Run go fmt against code.
.PHONY: fmt
fmt:
	@go fmt $(GO_PACKAGES)

test:
	go test ./... -v -cover