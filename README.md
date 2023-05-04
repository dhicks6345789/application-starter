# application-starter
A utility to help with the startup order of Windows applications.

## Who Is This Project For?
The code and settings contained in this project are intended for administrators of fleets of Microsoft Windows devices (laptops, desktop workstations, etc) who intend to have their users use the Google Drive client on their machines. This probably means administrators of school or company IT setups, the average home user is probably not going to find this very useful. In particular, you will probably need devices running an Enterprise or Education edition of Windows 10 or 11 (although all versions of Windows 11 might work, I'm not sure).

## What Does This Project Do?
<strong>The problem:</strong> when a user logs in to a Windows machine, the shell (Windows Explorer, the executable process which provides the familier Windows desktop and taskbar) loads. Then other applications load, including (if installed) Google Drive. Google Drive integrates with Windows Explorer to present a user's Google Drive as (for instance) "G:" drive, which is nice and easy for the user to work with.

By default, the user's Desktop folder will be a local (or domain) folder. A simple registry / Active Directory setting can redirect that location to a Google Drive folder ("G:\My Drive\Desktop", for instance), but as Explorer loads before Google Drive the user will get an error message pop-up saying their Desktop folder could not be found.

The solution to this is to make sure that Google Drive loads before Windows Explorer. This can be done by using a Windows registry entry to replace the defined user shell with an application that starts Google Drive, waits for it to be ready, then starts Explorer. This simple-sounding solution is complicated by some quirks of how Windows deals with some of the settings involved.

Replacing the shell application can be done by setting the "Shell" value in HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows NT\CurrentVersion\Winlogon. It's easy enough to create a simple executable that starts Google Drive, checks it's ready, then starts Explroer. However, the version of Explorer that starts will be the simple file folder view, not the full desktop - the user will be left with a file browser window instead of a proper desktop. Explorer only starts up as the desktop if that registry key is set to "Explorer.exe". Therefore, the solution is for the shell replacement application to start Google Drive, wait for it to start, set the registry key to "Explorer.exe", start Explorer, then set the registry key back to its own executable.

The next problem encountered is that the shell replacement executable starts as an unprivalged executable - it doesn't have permissions to go writing the appropriate registry keys. Therefore, we need a privalged helper process that runs as the system user and does have suitible permissions. This is all quite do-able, but does add to the complexity of the solution.

Install Chrome
Install Google Drive Client
Add auto-login if wanted
