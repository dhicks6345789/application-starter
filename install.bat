@echo off
echo Installing...

erase starter.exe
erase go.mod
copy application-starter\starter.go .
rem copy application-starter\go.mod .
rem go get golang.org/x/sys/windows/registry
go get github.com/luisiturrios/gowin
go build starter.go

if not exist "C:\Program Files\Application Starter" (
  mkdir "C:\Program Files\Application Starter"
)
copy /b/v/y starter.exe "C:\Program Files\Application Starter"
rem starter.exe

Set currentDir=%cd%
regedit /S %currentDir%\application-starter\settings.reg

echo Done!
