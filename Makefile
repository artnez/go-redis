all:
	@echo "Targets:"
	@echo "  make test"
	@echo "  make coverage"

test:
	go test -v ./...

coverage:
	go test -coverprofile cover.out ./...
	go tool cover -html=cover.out -o cover.html
	open cover.html
	sleep 1
	rm cover.{html,out}

.PHONY: coverage test
