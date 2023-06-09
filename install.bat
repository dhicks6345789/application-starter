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

echo Making sure Application Installer folder exists.
if not exist "C:\Program Files\Application Starter" (
  mkdir "C:\Program Files\Application Starter"
)

Set currentDir=%cd%

if exist application-starter (
  set subFolder="application-starter\"
  goto compileCode
)

if exist starter.go (
  set subFolder=""
  goto compileCode
)

goto downloadCode
:compileCode

echo Compiling starter.go...
go build application-starter\starter.go
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

regedit /S %currentDir%\application-starter\settings.reg
goto end

:downloadCode
echo Downloading starter.exe...
powershell -command "& {&'Invoke-WebRequest' -Uri https://www.sansay.co.uk/application-starter/binaries/starter-win-amd64.exe -OutFile 'C:\Program Files\Application Starter\starter.exe'}"
echo Downloading firstRun.exe...
powershell -command "& {&'Invoke-WebRequest' -Uri https://www.sansay.co.uk/application-starter/binaries/firstRun-win-amd64.exe -OutFile 'C:\Program Files\Application Starter\firstRun.exe'}"
echo Downloading setPerUser.reg...
powershell -command "& {&'Invoke-WebRequest' -Uri https://www.sansay.co.uk/application-starter/setPerUser.reg -OutFile 'C:\Program Files\Application Starter\setPerUser.reg'}"

echo Downloading and applying settings.reg...
powershell -command "& {&'Invoke-WebRequest' -Uri https://www.sansay.co.uk/application-starter/settings.reg -OutFile 'settings.reg'}"
regedit /S %currentDir%\settings.reg
erase settings.reg

:end
echo Done!
