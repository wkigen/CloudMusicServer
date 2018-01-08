xcopy %CD%\..\conf %CD%\conf /e /y
xcopy %CD%\..\conf %CD%\..\build\windows\conf /e /y

start ../build\windows\data_server.exe -log_dir="./log/" -alsologtostderr=true
start ../build\windows\login_server.exe -log_dir="./log/" -alsologtostderr=true
start ../build\windows\gate_server.exe -log_dir="./log/" -alsologtostderr=true
