SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64

go build -tags "reuseport quic kcp zookeeper etcd consul ping" -o ../build/linux/gate_server ../gateserver/gate_server_bin.go
go build -tags "reuseport quic kcp zookeeper etcd consul ping" -o ../build/linux/login_server ../loginserver/login_server_bin.go
go build -tags "reuseport quic kcp zookeeper etcd consul ping" -o ../build/linux/data_server ../dataserver/data_server_bin.go

pause