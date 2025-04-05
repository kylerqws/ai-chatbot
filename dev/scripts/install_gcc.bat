@echo off
chcp 65001 >nul
setlocal

:: Check if Chocolatey is installed
where choco >nul 2>nul
if %errorlevel%==0 (
    echo Installing MinGW-w64 via Chocolatey...
    choco install mingw -y
    echo MinGW-w64 installation via Chocolatey completed.
) else (
    echo Chocolatey not found. Trying to install via winget...

    :: Install MSYS2 via winget
    winget install -e --id MSYS2.MSYS2
    echo MSYS2 installation via winget completed.
)

:: Ensure GCC installation path
set "GCC_DIR=C:\msys64\mingw64\bin"

:: Convert to short (8.3) path format
for %%i in ("%GCC_DIR%") do set "GCC_DIR=%%~sfi"

:: Verify GCC installation (CHECK AGAIN)
if not exist "%GCC_DIR%\gcc.exe" (
    echo Error: GCC is not installed in "%GCC_DIR%"!
    echo Please check the installation manually.
    pause
    exit /b 1
)

:: Check if GCC directory is already in system PATH
echo %PATH% | findstr /I /C:"%GCC_DIR%" >nul
if %errorlevel% neq 0 (
    echo Adding GCC to system PATH...

    :: Append to system PATH
    setx PATH "%PATH%;%GCC_DIR%" >nul

    echo GCC has been added to PATH. Please restart your terminal.
) else (
    echo GCC is already in PATH.
)

pause
exit /b 0
