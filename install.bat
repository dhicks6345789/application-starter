@echo off
echo Installing...

erase starter.exe
copy application-starter\go.mod .
go get golang.org/x/sys/windows/registry
go build application-starter\starter.go

if not exist "C:\Program Files\Application Starter" (
  mkdir "C:\Program Files\Application Starter"
)
rem copy /b/v/y starter.exe "C:\Program Files\Application Starter"
rem starter.exe

Set currentDir=%cd%
regedit /S %currentDir%\application-starter\settings.reg

echo Done!
