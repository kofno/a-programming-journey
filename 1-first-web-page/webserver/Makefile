windows-server:
	GOOS=windows GOARCH=amd64 go build -o bin/windows/amd64/webserver.exe

linux-server:
	GOOS=linux GOARCH=amd64 go build -o bin/linux/amd64/webserver

mac-server:
	GOOS=darwin GOARCH=amd64 go build -o bin/mac/amd64/webserver
	GOOS=darwin GOARCH=arm64 go build -o bin/mac/arm64/webserver

all: windows-server linux-server mac-server
