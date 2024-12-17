go build -tags=tempdll -buildmode=exe -ldflags="-s -w -H windowsgui" -o bin/server.exe .

upx bin/server.exe

@REM pause