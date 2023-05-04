# application-starter
A utility to help with the startup order of Windows applications.

## Who Is This Project For?
The code and settings contained in this project are intended for administrators of fleets of Microsoft Windows devices (laptops, desktop workstations, etc) who intend to have their users use the Google Drive client on their machines. This probably means administrators of school or company IT setups, the average home user is probably not going to find this very useful. In particular, you will probably need devices running an Enterprise or Education edition of Windows 10 or 11 (although all versions of Windows 11 might work, I'm not sure).

## What Does This Project Do?
<strong>The problem:</strong> when a user logs in to a Windows machine, the shell (Windows Explorer, the executable process which provides the familier Windows desktop and taskbar) loads. Then other applications load, including (if installed) Google Drive. Google Drive integrates with Windows Explorer to present a user's Google Drive as (for instance) "G:" drive, which is nice and easy for the user to work with.

By default, the user's Desktop folder will be a local (or domain) folder. A simple registry / Active Directory setting can redirect that location to a Google Drive folder ("G:\My Drive\Desktop", for instance), but as Explorer loads before Google Drive the user will get an error message pop-up saying their Desktop folder could not be found.

The solution to this is to make sure that Google Drive loads before Windows Explorer. This can be done by using a Windows registry entry to replace the defined user shell with an application that starts Google Drive, waits for it to be ready, then starts Explorer.

Install Chrome
Install Google Drive Client
Add auto-login if wanted
