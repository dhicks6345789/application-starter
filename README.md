# Application Starter
A utility to help with the startup order of Windows applications - in particular, to make sure Google Drive starts before the Desktop so that the user desktop folder can be redirected to Google Drive.

## Who Is This Project For?
The code and settings contained in this project are intended for administrators of fleets of Microsoft Windows devices (laptops, desktop workstations, etc) who intend to have their users use the Google Drive client on their machines. This probably means administrators of school or company IT setups, the average home user is probably not going to find this very useful. In particular, you will probably need devices running an Enterprise or Education edition of Windows 10 or 11 (although all versions of Windows 11 might work, I'm not sure). Your organisation should probably also have a Google Workspace domain - I haven't tested this project with a personal account, although in theory there's no reason why people couldn't use personal accounts instead.

Combining this project with the [Google Drive client for Windows](https://www.google.com/intl/en-GB/drive/download/) and [Google Credential Provider for Windows](https://support.google.com/a/answer/9250996?hl=en) should provide a mechansim to have both cloud (i.e. Google) based authentication and file storage for a Windows machine without the need for an Active Directory setup.

## What Does This Project Do?

### Problem
When a user logs in to a Windows machine, the shell (Windows Explorer, the executable process which provides the familier Windows desktop and taskbar) loads. Then other applications load, including (if installed) Google Drive. Google Drive integrates with Windows Explorer to present a user's Google Drive as (for instance) "G:" drive, which is nice and easy for the user to work with.

By default, the user's Desktop folder will still be a local (or domain) folder. A simple registry / Active Directory setting can redirect that location to a Google Drive folder ("G:\My Drive\Desktop", for instance), but as Explorer loads before Google Drive the user will get an error message pop-up saying their Desktop folder could not be found.

The solution to this is to make sure that Google Drive loads before Windows Explorer. In theory, this could be done by using a Windows registry entry (the `Shell` value in `HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows NT\CurrentVersion\Winlogon`) to replace the defined user shell with an application that starts Google Drive, waits for it to be ready, then starts Explorer. However, the version of Explorer that starts up will be the simple file folder view, not the full desktop. Explorer only starts up as the desktop if that registry key is set to `Explorer.exe`.

Therefore, we run the helper application before Explorer starts by using the `Userinit` key instead. The helper application then stops Explorer, starts Google Drive, waits for it to be ready, then re-starts Explorer. A slightly different helper application is needed for first user login (as simply stopping Explorer stops the new-user setup process from completing), that is run using an entry in `HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows\CurrentVersion\Run`.

### Solution
This project contains a batch file installer and executable code, written in Go, to implement the setup as described above. It implements a small helper executable that stops Windows Explorer, starts the Google Drive client, waits for the Google Drive client to be ready, then starts Windows Explorer back up again.

## Installation
Before you install the code from this project you need to install the [Google Drive client](https://www.google.com/intl/en-GB/drive/download/) on your Windows desktop / laptop machine. You will also need to have [Google Chrome](https://www.google.com/intl/en_uk/chrome/) installed.

You don't need to sign in to the Google Drive client to run the installer. The installer will set a Registry option to tell the Drive client to bypass the system default browser option and always use Chrome.

The installer will also disable the Google Drive client start-on-login option as Google Drive is started separatly by the helper application. If the end user changes the option it shouldn't be any great problem, the Drive client will just pop up a status / file browser window on startup, just have the user change the option back again if they find that annoying.

The installer will set the user Registry entry to redirect the user's Desktop folder. It will also set the local desktop folder (`C:\Users\username\Desktop`) to be read-only to avoid the user trying to write anything there.

When a user logs in to a Windows computer for the first time they will be presented with the standard authorise-Google-Drive-for-this-machine dialog. They won't be able to proceed until they get Google Drive set up. If this is the first time a new user has logged in to any machine they will have a "Desktop" folder created in their "My Drive" section. Logging on to other machines, they should see the exact same desktop contents (although, typically, Windows re-arranges items on different desktops, especially if the screen resolution is different).

This project has been tested with the [Google Credential Provider for Windows](https://support.google.com/a/answer/9250996?hl=en) - a user should be able to log into Windows with their Google credentials and straight away be presented with the authorisation dialog to allow Google Drive access as well, credentials from the login should be passed through to Chrome.

Again, before you install anything, this project is intended for administrators of fleets of devices - this project is something you install on a test device of some kind, not on your or someone else's mission-critical personal workstation. It doesn't carry out any mass moving or deleting of files, but it does change the way your computer starts up, so you might be left with a blank screen on login (if you do find that, hitting ctrl-alt-delete should bring up a dialog that allows you to get a desktop environment going).

### Simple Install: One-Line Command
Open a command prompt as administrator and run the following line:

```
powershell -command "& {&'Invoke-WebRequest' -Uri https://www.sansay.co.uk/application-starter/install.bat -OutFile install.bat}" && install.bat && erase install.bat
```

Hopefully, that should be it - restart the machine and log in, you should be asked to set up your user account with Google Drive if it isn't already and your desktop should be redirected to `G:\My Drive\Desktop`.

### Less Simple Install: Download And Compile Source Code
Again, you'll need to open a command prompt as administrator. Note that the administrator command prompt starts in `C:\Windows\System32`, as you're going to be downloading files you should probably change folder (`cd C:\Users\admin` or whatever).

You will need the [Git](https://gitforwindows.org/) version control system and the [Go](https://go.dev/) programming language installed on your Windows machine.

Use Git to clone the source code:

```
git clone https://github.com/dhicks6345789/application-starter.git
```

This should result in a folder called "application-starter" in the current folder. Now, you can just run the install script:

```
application-starter\install.bat
```

With the source code present the script will compile the Go applications into executables rather than downloading them.

## Help / Support
This project comes with no garuntee of any further help or support, or even that it won't simply break your computer. I'm a systems administrator for a school, the code and setup in this project has been tested with the particular setup available to me (Windows 11 laptop, Google Workspace for Education). If you've found this project there's a good chance you are also the administrator for a school or company setup involving both Windows machines and a Google Workspace domain. If you have useful testing feedback or any suggestions you can open an issue on the [Github project](https://github.com/dhicks6345789/application-starter) or contact me over on [EduGeek](http://www.edugeek.net/members/dhicks.html) - if you're setting up Google Workspace / Chromebooks in an educational setting then EduGeek is a good community to browse and ask questions of.

A combination that hasn't been tested yet is this project working on a remote dekstop server - possibly one implemented by [another project of mine](https://github.com/dhicks6345789/remote-gateway). In theory, it should work okay - login cookies from the login process won't be passed through to Chrome like they are with GCPW, but otherwise the functionality should work the same. I plan to test this on a server soon.

### To Do
- More testing.
- Add restart loop for Google Drive setup for first use.
- Test on remote dekstop machine.
- Make list of startup options configurable to allow for different applications. Might be particularly handy as a way of setting up kisok machines for exams or displays.
- Add more redirects for Documents, check if there's other locations to set Registry keys so that all applications see redirected location properly.
