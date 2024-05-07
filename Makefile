run:
	go run ./cmd/main.go

swag:
	swag init -g cmd/api/main.go -o ./cmd/docs