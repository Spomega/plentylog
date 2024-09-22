.PHONY: lint
lint:
	golint  -set_exit_status $$(go list ./...)

.PHONY: test
test:
	go test -v ./...