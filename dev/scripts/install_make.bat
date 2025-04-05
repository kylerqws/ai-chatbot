@echo off
chcp 65001 >nul
setlocal

:: Check if Chocolatey is installed
where choco >nul 2>nul
if %errorlevel%==0 (
    echo Installing Make via Chocolatey...
    choco install make -y
    echo Make installation via Chocolatey completed.
) else (
    echo Chocolatey not found. Trying to install via winget...

    :: Install Make via winget
    winget install -e --id GnuWin32.Make
    echo Make installation via winget completed.
)

:: Ensure Make installation path
set "MAKE_DIR=C:\Program Files (x86)\GnuWin32\bin"

:: Convert to short (8.3) path format
for %%i in ("%MAKE_DIR%") do set "MAKE_DIR=%%~sfi"

:: Verify Make installation (CHECK AGAIN)
if not exist "%MAKE_DIR%\make.exe" (
    echo Error: Make is not installed in "%MAKE_DIR%"!
    echo Please check the installation manually.
    pause
    exit /b 1
)

:: Check if Make directory is already in system PATH
echo %PATH% | findstr /I /C:"%MAKE_DIR%" >nul
if %errorlevel% neq 0 (
    echo Adding Make to system PATH...

    :: Append to system PATH
    setx PATH "%PATH%;%MAKE_DIR%" >nul

    echo Make has been added to PATH. Please restart your terminal.
) else (
    echo Make is already in PATH.
)

pause
exit /b 0
