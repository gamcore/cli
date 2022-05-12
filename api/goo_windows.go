//go:build windows

package api

func shellScript() string {
	data := `@ECHO OFF
if exist "%GOO_PATH%/apps/goo/new {
	tasklist /fi "ImageName eq goo.exe" /fo csv 2>NUL | find /I "goo.exe">NUL
	if "%ERRORLEVEL%"=="0" taskkill /IM "goo.exe" /F
	ren "$GOO_PATH/apps/goo/current" "$GOO_PATH/apps/goo/old"
	ren "$GOO_PATH/apps/goo/new" "$GOO_PATH/apps/goo/current"
	rmdir /S "$GOO_PATH/apps/goo/old"
}

start $GOO_PATH/apps/goo/current/goo %*
`

	return data
}

func (u AppUpdateSchema) GetPattern() string {
	return u.Pattern.Windows
}
