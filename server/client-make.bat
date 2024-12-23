go-winres make --in="client/winres/winres.json" --out="client/winappres/rsrc"

go build -ldflags="-s -w" -o bin/client.exe ./client/main.go

upx bin/client.exe

pause