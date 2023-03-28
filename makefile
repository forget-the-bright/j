VERSION := 1.0.0
build_linux:
	set CGO_ENABLED=0
	set GOARCH=amd64
	set GOOS=linux
	go build -ldflags "-s -w -X main.version=${VERSION}" -o ./bin/linux/j
build_windows:
	set CGO_ENABLED=0
	set GOARCH=amd64
	set GOOS=windows
	go build -ldflags "-s -w -X main.version=${VERSION}" -o ./bin/windows/j.exe