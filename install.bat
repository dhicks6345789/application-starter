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

if exist C:\Users\d.hicks_knightsbridg (
  echo Deleting user...
  net user d.hicks_knightsbridg /delete
  del /Q /S C:\Users\d.hicks_knightsbridg
  rmdir /Q /S C:\Users\d.hicks_knightsbridg
)

if exist %userprofile%\AppData\Local\ApplicationStarter\starter.txt (
  del /Q /F %userprofile%\AppData\Local\ApplicationStarter\starter.txt 2>&1
)

echo Making sure Application Installer folder exists.
if not exist "C:\Program Files\Application Starter" (
  mkdir "C:\Program Files\Application Starter"
)

echo Compiling Go code...
go build -ldflags "-H windowsgui" application-starter\starter.go
rem go build -ldflags "-X main.debugOn=true" application-starter\starter.go
if not exist starter.exe (
  echo Compile fail - starter.go
  exit /B 1
)
copy /y starter.exe "C:\Program Files\Application Starter"
erase starter.exe

rem go build -ldflags "-H windowsgui" application-starter\firstRun.go
go build -ldflags "-X main.debugOn=true" application-starter\firstRun.go
if not exist firstRun.exe (
  echo Compile fail - firstRun.go
  exit /B 1
)
rem copy /y firstRun.exe "C:\ProgramData\Microsoft\Windows\Start Menu\Programs\Startup"
erase firstRun.exe

rem go build application-starter\service.go
rem if not exist service.exe (
  rem echo Compile fail - service.go
  rem exit /B 1
rem )

rem echo Stopping existing service...
rem net stop ApplicationStarter


rem if %debug%==1 (
rem   del /S /Q "C:\Program Files\Application Starter\Users\*"
rem )

rem copy /y service.exe "C:\Program Files\Application Starter"
rem erase service.exe
copy /y application-starter\setExplorer.reg "C:\Program Files\Application Starter"
copy /y application-starter\setStarter.reg "C:\Program Files\Application Starter"
copy /y application-starter\setPerUser.reg "C:\Program Files\Application Starter"

rem echo Setting up Windows service...
rem application-starter\nssm\2.24\win64\nssm install ApplicationStarter "C:\Program Files\Application Starter\service.exe" > nul 2>&1
rem application-starter\nssm\2.24\win64\nssm set ApplicationStarter Description "A shell replacement for Explorer that starts up Google Drive then Explorer, letting users be able to redirect their desktop folders to Google Drive." > nul 2>&1
rem application-starter\nssm\2.24\win64\nssm set ApplicationStarter DisplayName "Application Starter" > nul 2>&1
rem application-starter\nssm\2.24\win64\nssm set ApplicationStarter AppNoConsole 1 > nul 2>&1
rem application-starter\nssm\2.24\win64\nssm set ApplicationStarter Start SERVICE_AUTO_START > nul 2>&1

rem echo Starting service...
rem net start ApplicationStarter

Set currentDir=%cd%
regedit /S %currentDir%\application-starter\settings.reg

rem if %debug%==1 (
rem   echo Running Application Starter in debug mode...
rem   net stop ApplicationStarter
rem   "C:\Program Files\Application Starter\service.exe" --debug
rem )

echo Done!
