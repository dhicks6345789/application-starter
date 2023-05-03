@echo off

echo Building Go applications...
go build application-starter\starter.go
go build application-starter\service.go

echo Installing...
if not exist "C:\Program Files\Application Starter" (
  mkdir "C:\Program Files\Application Starter"
)
copy /b/v/y starter.exe "C:\Program Files\Application Starter"
erase starter.exe
copy /b/v/y service.exe "C:\Program Files\Application Starter"
erase service.exe
copy /b/v/y application-starter\setExplorer.reg "C:\Program Files\Application Starter"
copy /b/v/y application-starter\setStarter.reg "C:\Program Files\Application Starter"

echo Setting up Windows service...
application-starter\nssm-2.24\win64\nssm install ApplicationStarter "C:\Program Files\Application Starter\service.exe" > nul 2>&1
application-starter\nssm-2.24\win64\nssm set ApplicationStarter DisplayName "Application Starter" > nul 2>&1
application-starter\nssm-2.24\win64\nssm set ApplicationStarter AppNoConsole 1 > nul 2>&1
application-starter\nssm-2.24\win64\nssm set ApplicationStarter Start SERVICE_AUTO_START > nul 2>&1
net start ApplicationStarter

Set currentDir=%cd%
regedit /S %currentDir%\application-starter\settings.reg

echo Done!
