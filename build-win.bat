SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build -o ./bin/log-server.exe ./cmd/main.go 