@echo off
echo Installing...

erase starter.exe
go build application-starter/starter.go
starter.exe

echo Done!
