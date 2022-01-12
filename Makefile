run:
	go run main.go

install:
	go build -o bin/hilo main.go
	install -Dm755 bin/hilo /usr/bin/hilo
	
build:
	go build -o bin/hilo main.go