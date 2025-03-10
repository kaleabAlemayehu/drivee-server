.PHONY: all 

init:
	@docker compose up -d && goose up
.PHONY: init

dev: init
	@go run main.go 
.PHONY: dev

clean:
	@goose down && docker compose down
.PHONY: clean

test:
	@go test -v ./..
.PHONY: test
