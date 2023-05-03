@echo off
echo Installing...

erase starter.exe
cp application-starter\go.mod .
go install golang.org/x/sys/windows/registry@latest
go build application-starter\starter.go

if not exist "C:\Program Files\Application Starter" (
  mkdir "C:\Program Files\Application Starter"
)
copy /b/v/y starter.exe "C:\Program Files\Application Starter"
rem starter.exe

Set currentDir=%cd%
regedit /S %currentDir%\application-starter\settings.reg

echo Done!
