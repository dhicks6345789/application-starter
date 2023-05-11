@echo off
setlocal EnableDelayedExpansion

echo Installing Application Starter...

set debug=0

rem Parse any parameters.
:paramLoop
if "%1"=="" goto paramContinue
if "%1"=="--debug" (
  set debug=1
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

echo Compiling starter.go...
go build -ldflags application-starter\starter.go
if not exist starter.exe (
  echo Compile fail - starter.go
  exit /B 1
)
copy /y starter.exe "C:\Program Files\Application Starter"
erase starter.exe

echo Compiling firstRun.go...
go build -ldflags "-H windowsgui" application-starter\firstRun.go
if not exist firstRun.exe (
  echo Compile fail - firstRun.go
  exit /B 1
)
copy /y firstRun.exe "C:\Program Files\Application Starter"
erase firstRun.exe

copy /y application-starter\setPerUser.reg "C:\Program Files\Application Starter"

Set currentDir=%cd%
regedit /S %currentDir%\application-starter\settings.reg

echo Done!
