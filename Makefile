.PHONY: all 

init:
	@sudo docker compose up -d && goose -dir ./schema/migrations up && goose -dir ./schema/seeds -no-versioning up
.PHONY: init

dev: init
	@go run main.go 
.PHONY: dev

rerun:
	@go run main.go 
.PHONY: rerun 

populate:
	@goose -dir ./schema/migrations up && goose -dir ./schema/seeds -no-versioning up
.PHONY: populate

cleandb:
	@goose -dir ./schema/seeds -no-versioning reset && goose -dir ./schema/migrations down 
.PHONY: cleandb

clean:
	@goose -dir ./schema/seeds -no-versioning reset && goose -dir ./schema/migrations down && sudo docker compose down
.PHONY: clean

test:
	@go test -v ./..
.PHONY: test
