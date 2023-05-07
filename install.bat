@echo off
setlocal EnableDelayedExpansion

echo Installing Application Starter...

set debug=0

rem Parse any parameters.
:paramLoop
if "%1"=="" goto paramContinue
if "%1"=="--debug" (
  set debug=1
  echo ### DEBUG MODE SET ###
)
shift
goto paramLoop
:paramContinue

echo Compiling Go code...
if %debug%==1 (
  go build -ldflags "-X main.debugOn=true" application-starter\starter.go
) else (
  go build -ldflags "-H windowsgui" application-starter\starter.go
)
if not exist starter.exe (
  echo Compile fail - starter.go
  exit /B 1
)

go build application-starter\service.go
if not exist service.exe (
  echo Compile fail - service.go
  exit /B 1
)

echo Stopping existing service...
net stop ApplicationStarter

echo Installing...
if not exist "C:\Program Files\Application Starter" (
  mkdir "C:\Program Files\Application Starter"
  mkdir "C:\Program Files\Application Starter\Users"
)
if %debug%==1 (
  del /S /Q "C:\Program Files\Application Starter\Users\*"
)

copy /y starter.exe "C:\Program Files\Application Starter"
copy /y starter.exe "C:\Users\exams\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Satrtup"
erase starter.exe
copy /y service.exe "C:\Program Files\Application Starter"
erase service.exe
copy /y application-starter\setExplorer.reg "C:\Program Files\Application Starter"
copy /y application-starter\setStarter.reg "C:\Program Files\Application Starter"
copy /y application-starter\setPerUser.reg "C:\Program Files\Application Starter"

echo Setting up Windows service...
application-starter\nssm\2.24\win64\nssm install ApplicationStarter "C:\Program Files\Application Starter\service.exe" > nul 2>&1
application-starter\nssm\2.24\win64\nssm set ApplicationStarter Description "A shell replacement for Explorer that starts up Google Drive then Explorer, letting users be able to redirect their desktop folders to Google Drive." > nul 2>&1
application-starter\nssm\2.24\win64\nssm set ApplicationStarter DisplayName "Application Starter" > nul 2>&1
application-starter\nssm\2.24\win64\nssm set ApplicationStarter AppNoConsole 1 > nul 2>&1
application-starter\nssm\2.24\win64\nssm set ApplicationStarter Start SERVICE_AUTO_START > nul 2>&1

echo Starting service...
net start ApplicationStarter

Set currentDir=%cd%
regedit /S %currentDir%\application-starter\settings.reg

if %debug%==1 (
  echo Running Application Starter in debug mode...
  net stop ApplicationStarter
  "C:\Program Files\Application Starter\service.exe" --debug
)

echo Done!
