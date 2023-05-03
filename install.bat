@echo off
echo Installing...

erase starter.exe
go build application-starter\starter.go
starter.exe

regedit /S application-starter\settings.reg

echo Done!
