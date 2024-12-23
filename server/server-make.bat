go-winres make --in="host/winres/winres.json" --out="host/winappres/rsrc"

go build -tags=tempdll -buildmode=exe -ldflags="-s -w -H windowsgui" -o bin/server.exe ./host/main.go

upx bin/server.exe

pause