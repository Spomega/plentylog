.PHONY: lint
lint:
	golint  -set_exit_status $$(go list ./...)