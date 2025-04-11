.PHONY: all 

init:
	@sudo docker compose up -d && goose -dir ./schema/migrations up && goose -dir ./schema/seeds -no-versioning up
.PHONY: init

dev: init
	@go run main.go 
.PHONY: dev

clean:
	@goose -dir ./schema/seeds -no-versioning reset && goose -dir ./schema/migrations down && docker compose down
.PHONY: clean

test:
	@go test -v ./..
.PHONY: test
