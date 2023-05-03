@echo off
echo Installing...

erase starter.exe
go build application-starter/startyer.go
starter.exe

echo Done!
