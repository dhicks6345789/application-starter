@echo off
echo Installing...

erase starter.exe
go build application-starter\starter.go
starter.exe

Set currentDir=%cd%
regedit /S %currentDir%\application-starter\settings.reg

echo Done!
