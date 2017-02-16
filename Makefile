all: main.go
	go build -o webls main.go

linux: main.go
	env GOOS=linux GOARCH=amd64 go build -o webls-linux main.go

clean:
	rm webls