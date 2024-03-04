@REM This runs the x86_64 version of the CI binary

SET ROOT_DIR=%cd%
.\.github\scripts\dist\yacar_ci_windows_amd64.exe
PAUSE
