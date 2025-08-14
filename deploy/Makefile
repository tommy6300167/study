APP_NAME=you-mon-so-fat

run:
	go run main.go

build:
	go build -o main .

docker-build:
	docker build -t $(APP_NAME) .

docker-run:
	docker run --env-file .env -p 8080:8080 $(APP_NAME)

test:
	go test ./...

clean:
	rm -f main