# application-starter
A utility to help with the startup order of Windows applications - in particular, to make sure Google Drive startes before the Desktop so that the user desktop folder can be redirected to Google Drive.

## Who Is This Project For?
The code and settings contained in this project are intended for administrators of fleets of Microsoft Windows devices (laptops, desktop workstations, etc) who intend to have their users use the Google Drive client on their machines. This probably means administrators of school or company IT setups, the average home user is probably not going to find this very useful. In particular, you will probably need devices running an Enterprise or Education edition of Windows 10 or 11 (although all versions of Windows 11 might work, I'm not sure).

## What Does This Project Do?

### Problem
When a user logs in to a Windows machine, the shell (Windows Explorer, the executable process which provides the familier Windows desktop and taskbar) loads. Then other applications load, including (if installed) Google Drive. Google Drive integrates with Windows Explorer to present a user's Google Drive as (for instance) "G:" drive, which is nice and easy for the user to work with.

By default, the user's Desktop folder will still be a local (or domain) folder. A simple registry / Active Directory setting can redirect that location to a Google Drive folder ("G:\My Drive\Desktop", for instance), but as Explorer loads before Google Drive the user will get an error message pop-up saying their Desktop folder could not be found.

The solution to this is to make sure that Google Drive loads before Windows Explorer. In theory, this could be done by using a Windows registry entry (the `Shell` value in `HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows NT\CurrentVersion\Winlogon`) to replace the defined user shell with an application that starts Google Drive, waits for it to be ready, then starts Explorer. However, the version of Explorer that starts up will be the simple file folder view, not the full desktop. Explorer only starts up as the desktop if that registry key is set to `Explorer.exe`.

Therefore, we run the helper application before Explorer starts by using the `Userinit` key instead. The helper application then stops Explorer, starts Google Drive, waits for it to start, then re-starts Explorer. A slightly different helper application is needed for first user login (as simply stopping Explorer stops the new-user setup process from completing), that is run using an entry in `HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows\CurrentVersion\Run`.

### Solution
This project contains a batch file installer and executable code, written in Go, to implement the setup as described above.

## Installation
### One-Line Command
Open a command prompt as administrator and run the following line:

`powershell -command "& {&'Invoke-WebRequest' -Uri https://www.sansay.co.uk/application-starter/install.bat -OutFile install.bat}" && install.bat && erase install.bat`

### Download Source Code
Again, you'll need to open a command prompt as administrator. Note that the administrator command prompt starts in C:\Windows\System32, as you're going to be downloading files you should probably change folder (`cd C:\Users\admin` or whatever).

You'll need [Git](https://gitforwindows.org/) installed on your Windows machine. Clone the project:

`git clone https://github.com/dhicks6345789/application-starter.git`

This should result in a folder called "application-starter" in the current folder.

## To Do
- Have batch file download and runnable in one line.
- Loading screen for new user section.
- Check all folder redirects are correct.
