rsrc -manifest main.go.manifest -o rsrc.syso
GOOS=windows GOARCH=amd64 go build -ldflags="-H windowsgui" 
