@echo off
echo Installing...

erase starter.exe
copy application-starter\starter.go .
copy application-starter\go.mod .
rem go get golang.org/x/sys/windows/registry
go install github.com/luisiturrios/gowin@latest
go build starter.go

if not exist "C:\Program Files\Application Starter" (
  mkdir "C:\Program Files\Application Starter"
)
copy /b/v/y starter.exe "C:\Program Files\Application Starter"
rem starter.exe

Set currentDir=%cd%
regedit /S %currentDir%\application-starter\settings.reg

echo Done!
