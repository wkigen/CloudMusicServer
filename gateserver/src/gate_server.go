package gateserver

import (
	"flag"
	"../../log"
	"github.com/smallnest/rpcx/client"
)

var (
	addr       = flag.String("addr", "localhost:8701", "http server address")
	zkAddr     = flag.String("zkAddr", "localhost:2181", "zookeeper address")
	st         = flag.String("st", "http1", "server type: http1 or h2c")
	basePath   = flag.String("basepath", "/cloudmusic", "basepath for zookeeper, etcd and consul")
	failmode   = flag.Int("failmode", int(client.Failtry), "failMode, Failover in default")
	selectMode = flag.Int("selectmode", int(client.RandomSelect), "selectMode, RoundRobin in default")
)

func Start() {
	flag.Parse()

	gw := NewGateway(*basePath,*addr,*zkAddr,  ServerType(*st), client.FailMode(*failmode), client.SelectMode(*selectMode), client.DefaultOption)
	log.Log(log.Info,"start gate server")
	gw.Serve()
}