go build -tags=tempdll -buildmode=exe -ldflags="-s -w -H windowsgui" -o bin/server.exe ./cmd/host/main.go

go build -ldflags="-s -w" -o bin/client.exe ./cmd/client/main.go

upx bin/server.exe
upx bin/client.exe

pause