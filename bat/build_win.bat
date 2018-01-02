go build -tags "reuseport quic kcp zookeeper etcd consul ping" -o ../build/windows/gate_server.exe ../gateserver/gate_server_bin.go
go build -tags "reuseport quic kcp zookeeper etcd consul ping" -o ../build/windows/login_server.exe ../loginserver/login_server_bin.go

pause