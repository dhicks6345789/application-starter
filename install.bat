@echo off
echo Installing...

erase starter.exe
go build application-starter\starter.go

if not exist "C:\Program Files\Application Starter" (
  mkdir "C:\Program Files\Application Starter"
)
copy /b/v/y starter.exe "C:\Program Files\Application Starter"
copy /b/v/y application-starter\setExplorer.reg "C:\Program Files\Application Starter"
copy /b/v/y application-starter\setStarter.reg "C:\Program Files\Application Starter"

Set currentDir=%cd%
regedit /S %currentDir%\application-starter\settings.reg

echo Done!
