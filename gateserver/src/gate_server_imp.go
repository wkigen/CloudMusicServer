package gateserver

import (
	"fmt"
	"context"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/rpcx/client"
	"golang.org/x/net/http2"
)

type ServerType string

const (
	HTTP1  ServerType = "http1"
	HTTP2             = "http2"
	HTTP2c            = "h2c"
)

type Gateway struct {
	BasePath       	string
	// http listen address
	Addr       		string
	ZookeeperAddr 	[]string
	ServerToken		string
	ServerType ServerType

	serviceDiscovery client.ServiceDiscovery
	FailMode         client.FailMode
	SelectMode       client.SelectMode
	Option           client.Option

	mu       sync.RWMutex
	xclients map[string]client.XClient
}

func NewGateway(basePath string,token string,addr string, zkAddr []string,st ServerType, failMode client.FailMode, selectMode client.SelectMode, option client.Option) *Gateway {
	return &Gateway{
		BasePath:		  basePath,
		Addr:             addr,
		ZookeeperAddr:	  zkAddr,
		ServerToken:	  token,
		ServerType:       st,
		FailMode:         failMode,
		SelectMode:       selectMode,
		Option:           option,
		xclients:         make(map[string]client.XClient),
	}
}

func (g *Gateway) Serve() {
	router := httprouter.New()
	router.POST("/*servicePath", g.handleRequest)
	router.GET("/*servicePath", g.handleRequest)
	router.PUT("/*servicePath", g.handleRequest)

	switch g.ServerType {
	case HTTP2c:
		g.startH2c(router)
	case HTTP2:
		panic("unsupported")
	default:
		g.startHttp1(router)
	}
}

func (g *Gateway) startHttp1(handler http.Handler) {
	if err := http.ListenAndServe(g.Addr, handler); err != nil {
		glog.Fatal("error in ListenAndServe: %s", err)
	}
}

func (g *Gateway) startH2c(handler http.Handler) {
	server := &http.Server{
		Addr:         g.Addr,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	//http2.Server.ServeConn()
	s2 := &http2.Server{
		IdleTimeout: 1 * time.Minute,
	}
	http2.ConfigureServer(server, s2)
	l, _ := net.Listen("tcp", g.Addr)
	defer l.Close()
	glog.Infoln("Start server...")
	for {
		rwc, err := l.Accept()
		if err != nil {
			glog.Errorln("accept err:", err)
			continue
		}
		go s2.ServeConn(rwc, &http2.ServeConnOpts{BaseConfig: server})
	}
}

func (g *Gateway) Interceptor(servicePath string) (error){

	if(servicePath == "DataServer"){
		return fmt.Errorf("can non connect to the service ,%s",servicePath)
	}

	return nil
}

func (g *Gateway) handleRequest(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if r.Header.Get(XServicePath) == "" {
		servicePath := params.ByName("servicePath")
		if strings.HasPrefix(servicePath, "/") {
			servicePath = servicePath[1:]
		}

		r.Header.Set(XServicePath, servicePath)
	}

	servicePath := r.Header.Get(XServicePath)

	wh := w.Header()

	err := g.Interceptor(servicePath)
	if(err != nil){
		glog.Warningln(err)
		wh.Set(XMessageStatusType, "Error")
		wh.Set(XErrorMessage, err.Error())
		w.Write([] byte("error"))
		return
	}

	req, err := HttpRequest2RpcxRequest(r,g.ServerToken)
	if err != nil {
		rh := r.Header
		for k, v := range rh {
			if strings.HasPrefix(k, "X-RPCX-") && len(v) > 0 {
				wh.Set(k, v[0])
			}
		}

		wh.Set(XMessageStatusType, "Error")
		wh.Set(XErrorMessage, err.Error())
		return
	}

	var xc client.XClient
	g.mu.Lock()
	if g.xclients[servicePath] == nil {
		zd := client.NewZookeeperDiscovery("/"+g.BasePath, servicePath, g.ZookeeperAddr, nil)
		g.xclients[servicePath] = client.NewXClient(servicePath, g.FailMode, g.SelectMode,zd, g.Option)
	}
	xc = g.xclients[servicePath]

	g.mu.Unlock()
	m, payload, err := xc.SendRaw(context.Background(), req)

	for k, v := range m {
		wh.Set(k, v)
	}
	if err != nil {
		wh.Set(XMessageStatusType, "Error")
		wh.Set(XErrorMessage, err.Error())
		return
	}

	w.Write(payload)

}
