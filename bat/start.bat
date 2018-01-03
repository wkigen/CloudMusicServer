xcopy %CD%\..\conf %CD%\conf /e /y
xcopy %CD%\..\conf %CD%\..\build\windows\conf /e /y

start ../build\windows\data_server.exe
start ../build\windows\login_server.exe
start ../build\windows\gate_server.exe
