
if not exist %CD%\log\data_server (
    md %CD%\log\data_server
)

if not exist %CD%\log\login_server (
    md %CD%\log\login_server
)

if not exist %CD%\log\gate_server (
    md %CD%\log\gate_server
)

REM if not exist %CD%\conf  (
REM     md %CD%\conf 
REM )

REM if not exist %CD%\..\build\windows\conf (
REM     md %CD%\..\build\windows\conf
REM )

REM xcopy %CD%\..\conf %CD%\conf /e /y
REM xcopy %CD%\..\conf %CD%\..\build\windows\conf /e /y

start ../build\windows\data_server.exe -log_dir="./log/data_server/" -alsologtostderr=true -config="../conf/config.xml"
start ../build\windows\login_server.exe -log_dir="./log/login_server/" -alsologtostderr=true -config="../conf/config.xml"
start ../build\windows\gate_server.exe -log_dir="./log/gate_server/" -alsologtostderr=true -config="../conf/config.xml"
